package validatex

import (
	"fmt"
)

type ArrayValidator struct {
	*Validator
	value []any
}

func NewArrayValidator(value []any) *ArrayValidator {
	return &ArrayValidator{
		Validator: new(Validator),
		value:     value,
	}
}
func (v *ArrayValidator) MaxLen(n int) *ArrayValidator {
	if v.Error != nil {
		return v
	}
	if len(v.value) > n {
		v.Error = fmt.Errorf("%w: maximum length is %d", ErrMaxLenCheckFailed, n)
	}
	return v
}

func (v *ArrayValidator) MinLen(n int) *ArrayValidator {
	if v.Error != nil {
		return v
	}
	if len(v.value) < n {
		v.Error = fmt.Errorf("%w: minimum length is %d", ErrMinLenCheckFailed, n)
	}
	return v
}

func (v *ArrayValidator) Items(fn func(item any, validator *Validator)) *ArrayValidator {
	if v.Error != nil {
		return v
	}
	for _, item := range v.value {
		fn(item, v.Validator)
	}
	return v
}
