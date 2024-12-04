package taskx

const (
	DefaultNamespace = "taskx"
	KeySeparator     = ":"

	defaultWorkerSize        = 4
	defaultWorkerPool        = 10
	defaultHeartbeatTTL      = 60  // seconds
	defaultHeartbeatInterval = 5   // seconds
	defaultLockTimeout       = 300 // seconds
	defaultRetryDelay        = 100 // milliseconds
	defaultRetryCount        = 3
)
