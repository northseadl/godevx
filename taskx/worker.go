package taskx

import (
	"context"
	"runtime/debug"
	"time"
)

type Worker struct {
	id       string
	poolSize int
	tasks    chan Task
	tm       *TaskManager
	stopCh   chan struct{}
}

func NewWorker(id string, poolSize int, tm *TaskManager) *Worker {
	return &Worker{
		id:       id,
		poolSize: poolSize,
		tasks:    make(chan Task, poolSize),
		tm:       tm,
		stopCh:   make(chan struct{}),
	}
}

func (w *Worker) Start(ctx context.Context) {
	pool := make(chan struct{}, w.poolSize)

	go w.heartbeat(ctx)

	for {
		select {
		case <-ctx.Done():
			return
		case <-w.stopCh:
			return
		case task := <-w.tasks:
			select {
			case pool <- struct{}{}:
				go func(t Task) {
					defer func() {
						<-pool
					}()
					w.executeTask(ctx, t)
				}(task)
			default:
				// 工作池满，等待下一次调度
			}
		}
	}
}

func (w *Worker) Stop() {
	close(w.stopCh)
}

func (w *Worker) executeTask(ctx context.Context, task Task) {
	result := &TaskResult{
		TaskID:    task.GetID(),
		StartTime: time.Now(),
	}

	defer func() {
		result.EndTime = time.Now()
		result.Duration = result.EndTime.Sub(result.StartTime)

		if r := recover(); r != nil {
			result.Status = TaskStatusFailed
			result.PanicError = r
			result.StackTrace = debug.Stack()
			w.tm.triggerHooks(func(h TaskHook) error {
				return h.OnTaskPanic(task, result)
			})
		}
	}()

	// 获取任务锁
	lockKey := w.tm.keyManager.TaskLockKey(task.GetID())
	locked, err := w.tm.acquireLock(ctx, lockKey)
	if err != nil || !locked {
		result.Status = TaskStatusFailed
		result.Error = ErrTaskLockFailed
		return
	}

	// 执行任务
	w.tm.triggerHooks(func(h TaskHook) error {
		return h.OnTaskStart(task)
	})

	taskCtx, cancel := context.WithTimeout(ctx, task.GetConfig().(*BaseTaskConfig).Timeout)
	defer cancel()

	err = task.Execute(taskCtx)

	if err != nil {
		result.Status = TaskStatusFailed
		result.Error = err
		w.tm.triggerHooks(func(h TaskHook) error {
			return h.OnTaskFail(task, result)
		})
	} else {
		result.Status = TaskStatusCompleted
		w.tm.triggerHooks(func(h TaskHook) error {
			return h.OnTaskComplete(task, result)
		})
	}
}

func (w *Worker) heartbeat(ctx context.Context) {
	ticker := time.NewTicker(time.Second * defaultHeartbeatInterval)
	defer ticker.Stop()

	heartbeatKey := w.tm.keyManager.WorkerHeartbeatKey(w.id)

	for {
		select {
		case <-ctx.Done():
			return
		case <-w.stopCh:
			return
		case <-ticker.C:
			w.tm.redis.Set(ctx,
				heartbeatKey,
				time.Now().Unix(),
				time.Second*defaultHeartbeatTTL)
		}
	}
}
