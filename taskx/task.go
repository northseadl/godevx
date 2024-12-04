package taskx

import (
	"context"
	"time"
)

type Task interface {
	Execute(ctx context.Context) error
	GetID() string
	GetType() TaskType
	GetConfig() TaskConfig
}

type TaskConfig interface {
	Validate() error
}

type BaseTaskConfig struct {
	ID          string
	Description string
	Timeout     time.Duration
	RetryCount  int
	Tags        []string
}

func (c *BaseTaskConfig) Validate() error {
	if c.ID == "" {
		return ErrInvalidConfig
	}
	return nil
}

type ScheduleTaskConfig struct {
	BaseTaskConfig
	Cron string
}

type ContinuousTaskConfig struct {
	BaseTaskConfig
	Interval time.Duration
}

type OnceTaskConfig struct {
	BaseTaskConfig
	ExecuteAt time.Time
}
