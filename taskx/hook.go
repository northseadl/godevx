package taskx

// TaskHook 定义任务生命周期钩子
type TaskHook interface {
	OnTaskStart(task Task) error
	OnTaskComplete(task Task, result *TaskResult) error
	OnTaskFail(task Task, result *TaskResult) error
	OnTaskPanic(task Task, result *TaskResult) error
}

// NoopTaskHook 提供空实现
type NoopTaskHook struct{}

func (h *NoopTaskHook) OnTaskStart(task Task) error                        { return nil }
func (h *NoopTaskHook) OnTaskComplete(task Task, result *TaskResult) error { return nil }
func (h *NoopTaskHook) OnTaskFail(task Task, result *TaskResult) error     { return nil }
func (h *NoopTaskHook) OnTaskPanic(task Task, result *TaskResult) error    { return nil }
