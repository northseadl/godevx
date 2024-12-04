package taskx

import "errors"

var (
	ErrTaskNotFound   = errors.New("task not found")
	ErrTaskLockFailed = errors.New("failed to acquire task lock")
	ErrTaskTimeout    = errors.New("task execution timeout")
	ErrInvalidConfig  = errors.New("invalid task configuration")
	ErrWorkerStopped  = errors.New("worker has been stopped")
	ErrManagerStopped = errors.New("task manager has been stopped")
)
