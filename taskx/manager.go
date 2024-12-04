package taskx

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"

	"github.com/google/uuid"
)

type TaskManager struct {
	redis      *redis.Client
	keyManager *KeyManager
	tasks      map[string]Task
	hooks      []TaskHook
	workers    []*Worker
	workerSize int
	poolSize   int
	mu         sync.RWMutex
	ctx        context.Context
	cancel     context.CancelFunc
}

func NewTaskManager(redisClient *redis.Client, opts ...Option) *TaskManager {
	options := DefaultOptions()
	for _, opt := range opts {
		opt(&options)
	}

	ctx, cancel := context.WithCancel(context.Background())

	tm := &TaskManager{
		redis:      redisClient,
		keyManager: NewKeyManager(options.Namespace),
		tasks:      make(map[string]Task),
		hooks:      options.Hooks,
		workerSize: options.WorkerSize,
		poolSize:   options.PoolSize,
		ctx:        ctx,
		cancel:     cancel,
	}

	tm.initWorkers()

	return tm
}

func (tm *TaskManager) initWorkers() {
	tm.workers = make([]*Worker, tm.workerSize)
	for i := 0; i < tm.workerSize; i++ {
		tm.workers[i] = NewWorker(
			fmt.Sprintf("worker-%d-%s", i, uuid.New().String()),
			tm.poolSize,
			tm,
		)
	}
}

func (tm *TaskManager) Start() {
	for _, worker := range tm.workers {
		go worker.Start(tm.ctx)
	}

	go tm.dispatcher()
}

func (tm *TaskManager) Stop() {
	tm.cancel()
	for _, worker := range tm.workers {
		worker.Stop()
	}
}

func (tm *TaskManager) RegisterTask(task Task) error {
	if err := task.GetConfig().Validate(); err != nil {
		return err
	}

	tm.mu.Lock()
	defer tm.mu.Unlock()

	tm.tasks[task.GetID()] = task
	return nil
}

func (tm *TaskManager) triggerHooks(fn func(TaskHook) error) {
	for _, hook := range tm.hooks {
		if err := fn(hook); err != nil {
			// 可以考虑记录错误或进行其他处理
			continue
		}
	}
}

func (tm *TaskManager) acquireLock(ctx context.Context, key string) (bool, error) {
	return tm.redis.SetNX(ctx, key, "1",
		time.Second*defaultLockTimeout).Result()
}

func (tm *TaskManager) dispatcher() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-tm.ctx.Done():
			return
		case <-ticker.C:
			activeWorkers := tm.getActiveWorkers()
			if len(activeWorkers) == 0 {
				continue
			}

			tm.dispatchTasks(activeWorkers)
		}
	}
}

func (tm *TaskManager) getActiveWorkers() []*Worker {
	var activeWorkers []*Worker

	for _, worker := range tm.workers {
		heartbeatKey := tm.keyManager.WorkerHeartbeatKey(worker.id)
		ts, err := tm.redis.Get(tm.ctx, heartbeatKey).Int64()
		if err != nil {
			continue
		}

		if time.Now().Unix()-ts < defaultHeartbeatTTL {
			activeWorkers = append(activeWorkers, worker)
		}
	}

	return activeWorkers
}

func (tm *TaskManager) dispatchTasks(workers []*Worker) {
	queueKey := tm.keyManager.TaskQueueKey()
	taskIDs, err := tm.redis.LRange(tm.ctx, queueKey, 0,
		int64(len(workers))-1).Result()
	if err != nil || len(taskIDs) == 0 {
		return
	}

	for i, taskID := range taskIDs {
		tm.mu.RLock()
		task, exists := tm.tasks[taskID]
		tm.mu.RUnlock()

		if !exists {
			continue
		}

		workerIdx := i % len(workers)
		select {
		case workers[workerIdx].tasks <- task:
			tm.redis.LRem(tm.ctx, queueKey, 1, taskID)
		default:
			// Worker队列已满，等待下次调度
		}
	}
}
