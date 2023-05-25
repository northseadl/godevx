package validators

import (
	"fmt"
	"godevx/validatex"
)

type FloatValidator struct {
	*validatex.Validator
	value float64
}

func (v *FloatValidator) Max(maxValue float64) *FloatValidator {
	if v.Error != nil {
		return v
	}
	if v.value > maxValue {
		v.Error = fmt.Errorf("%w: maximum Value is %f", ErrMaxValueCheckFailed, maxValue)
	}
	return v
}

func (v *FloatValidator) Min(minValue float64) *FloatValidator {
	if v.Error != nil {
		return v
	}
	if v.value < minValue {
		v.Error = fmt.Errorf("%w: minimum Value is %f", ErrMinValueCheckFailed, minValue)
	}
	return v
}
