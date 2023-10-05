package convx

// Ptr 取指针
func Ptr[T any](value T) *T {
	if value == nil {
		return nil
	}
	return &value
}
