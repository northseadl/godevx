package convx

// Ptr 取指针
func Ptr[T any](value T) *T {
	return &value
}
