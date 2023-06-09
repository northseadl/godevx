package validators

import (
	"fmt"
	"github.com/northseadl/godevx/validatex"
)

type FloatValidator struct {
	*validatex.Validator
	Value float64
}

func NewFloatValidator(value float64) *FloatValidator {
	return &FloatValidator{
		Validator: new(validatex.Validator),
		Value:     value,
	}
}

func (v *FloatValidator) Max(maxValue float64) *FloatValidator {
	if v.Error != nil {
		return v
	}
	if v.Value > maxValue {
		v.Error = fmt.Errorf("%w: maximum Value is %f", ErrMaxValueCheckFailed, maxValue)
	}
	return v
}

func (v *FloatValidator) Min(minValue float64) *FloatValidator {
	if v.Error != nil {
		return v
	}
	if v.Value < minValue {
		v.Error = fmt.Errorf("%w: minimum Value is %f", ErrMinValueCheckFailed, minValue)
	}
	return v
}
