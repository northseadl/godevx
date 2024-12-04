package taskx

type Options struct {
	Namespace  string
	WorkerSize int
	PoolSize   int
	Hooks      []TaskHook
}

func DefaultOptions() Options {
	return Options{
		Namespace:  DefaultNamespace,
		WorkerSize: defaultWorkerSize,
		PoolSize:   defaultWorkerPool,
		Hooks:      []TaskHook{&NoopTaskHook{}},
	}
}

type Option func(*Options)

func WithNamespace(namespace string) Option {
	return func(o *Options) {
		o.Namespace = namespace
	}
}

func WithWorkerSize(size int) Option {
	return func(o *Options) {
		o.WorkerSize = size
	}
}

func WithPoolSize(size int) Option {
	return func(o *Options) {
		o.PoolSize = size
	}
}

func WithHooks(hooks ...TaskHook) Option {
	return func(o *Options) {
		o.Hooks = hooks
	}
}
