package taskx

import "time"

type TaskType int

const (
	TaskTypeSchedule TaskType = iota
	TaskTypeContinuous
	TaskTypeOnce
)

type TaskStatus int

const (
	TaskStatusPending TaskStatus = iota
	TaskStatusRunning
	TaskStatusCompleted
	TaskStatusFailed
	TaskStatusTimeout
)

// TaskResult 任务执行结果
type TaskResult struct {
	TaskID     string
	Status     TaskStatus
	StartTime  time.Time
	EndTime    time.Time
	Duration   time.Duration
	Error      error
	PanicError interface{}
	StackTrace []byte
}
