package validators

import (
	"fmt"
	"github.com/northseadl/godevx/validatex"
)

type ArrayValidator[T any] struct {
	*validatex.Validator
	value []T
}

func NewArrayValidator[T any](value []T) *ArrayValidator[T] {
	return &ArrayValidator[T]{
		Validator: new(validatex.Validator),
		value:     value,
	}
}

func (v *ArrayValidator[T]) MaxLen(n int) *ArrayValidator[T] {
	if v.Error != nil {
		return v
	}
	if len(v.value) > n {
		v.Error = fmt.Errorf("%w: maximum length is %d", ErrMaxLenCheckFailed, n)
	}
	return v
}

func (v *ArrayValidator[T]) MinLen(n int) *ArrayValidator[T] {
	if v.Error != nil {
		return v
	}
	if len(v.value) < n {
		v.Error = fmt.Errorf("%w: minimum length is %d", ErrMinLenCheckFailed, n)
	}
	return v
}

func (v *ArrayValidator[T]) Items(fn func(item T, validator *validatex.Validator)) {
	if v.Error != nil {
		return
	}
	for _, item := range v.value {
		fn(item, v.Validator)
	}
}
