# Taskx

Taskx 是一个强大的分布式任务调度框架，专为 Go 语言设计，支持多实例部署和多种任务类型。

## 特性

- 支持三种任务类型：
  - 定时任务（Schedule Task）
  - 持续任务（Continuous Task）
  - 单次任务（Once Task）
- 分布式协调与任务锁
- 多 Worker 并行处理
- Panic 恢复机制
- 完善的生命周期钩子
- 基于 Redis 的可靠存储
- 灵活的配置选项
- 优雅的性能控制
- 完整的可观测性支持

## 安装

```bash
go get github.com/yourusername/taskx
```

## 快速开始

### 基础使用

```go
package main

import (
    "context"
    "time"
    
    "github.com/go-redis/redis/v8"
    "github.com/yourusername/taskx"
)

// 定义任务
type MyTask struct {
    taskx.BaseTaskConfig
}

func (t *MyTask) Execute(ctx context.Context) error {
    // 实现任务逻辑
    return nil
}

func (t *MyTask) GetID() string {
    return t.ID
}

func (t *MyTask) GetType() taskx.TaskType {
    return taskx.TaskTypeSchedule
}

func (t *MyTask) GetConfig() taskx.TaskConfig {
    return &t.BaseTaskConfig
}

func main() {
    // 创建Redis客户端
    rdb := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })

    // 创建任务管理器
    tm := taskx.NewTaskManager(rdb,
        taskx.WithNamespace("myapp"),
        taskx.WithWorkerSize(4),
        taskx.WithPoolSize(10),
    )

    // 注册任务
    task := &MyTask{
        BaseTaskConfig: taskx.BaseTaskConfig{
            ID:          "task-1",
            Description: "My first task",
            Timeout:     time.Minute,
        },
    }
    
    tm.RegisterTask(task)

    // 启动任务管理器
    tm.Start()
    defer tm.Stop()

    // 等待信号退出
    select {}
}
```

### 自定义Hook

```go
type CustomHook struct {
    logger *log.Logger
}

func (h *CustomHook) OnTaskStart(task taskx.Task) error {
    h.logger.Printf("Task started: %s", task.GetID())
    return nil
}

func (h *CustomHook) OnTaskComplete(task taskx.Task, result *taskx.TaskResult) error {
    h.logger.Printf("Task completed: %s, duration: %v", task.GetID(), result.Duration)
    return nil
}

func (h *CustomHook) OnTaskFail(task taskx.Task, result *taskx.TaskResult) error {
    h.logger.Printf("Task failed: %s, error: %v", task.GetID(), result.Error)
    return nil
}

func (h *CustomHook) OnTaskPanic(task taskx.Task, result *taskx.TaskResult) error {
    h.logger.Printf("Task panic: %s, error: %v", task.GetID(), result.PanicError)
    return nil
}
```

## 配置选项

### TaskManager 选项

```go
// 设置命名空间
taskx.WithNamespace("myapp")

// 设置Worker数量
taskx.WithWorkerSize(4)

// 设置每个Worker的协程池大小
taskx.WithPoolSize(10)

// 添加Hook
taskx.WithHooks(customHook)
```

### 任务配置

```go
// 基础配置
type BaseTaskConfig struct {
    ID          string
    Description string
    Timeout     time.Duration
    RetryCount  int
    Tags        []string
}

// 定时任务配置
type ScheduleTaskConfig struct {
    BaseTaskConfig
    Cron string
}

// 持续任务配置
type ContinuousTaskConfig struct {
    BaseTaskConfig
    Interval time.Duration
}

// 单次任务配置
type OnceTaskConfig struct {
    BaseTaskConfig
    ExecuteAt time.Time
}
```

## 最佳实践

1. **命名空间管理**
    - 为不同环境使用不同的命名空间
    - 开发环境和生产环境分开配置

2. **Worker配置**
    - 根据机器配置和任务特点调整Worker数量
    - 合理设置协程池大小避免资源耗尽

3. **任务设计**
    - 合理设置任务超时时间
    - 实现幂等性设计
    - 做好任务拆分

4. **监控和告警**
    - 实现自定义Hook记录关键指标
    - 设置合理的告警阈值
    - 定期检查任务执行情况

## 性能优化建议

1. Redis连接池配置
```go
rdb := redis.NewClient(&redis.Options{
    Addr:         "localhost:6379",
    PoolSize:     10,
    MinIdleConns: 5,
    MaxRetries:   3,
})
```

2. 任务批量处理
```go
// 实现批量任务接口
type BatchTask interface {
    Task
    BatchSize() int
}
```

3. 合理的任务分片
```go
// 任务分片接口
type ShardingTask interface {
    Task
    GetShardingKey() string
}
```