package validatex

import (
	"fmt"
)

type FloatValidator struct {
	*Validator
	Value float64
}

func NewFloatValidator(value float64) *FloatValidator {
	return &FloatValidator{
		Validator: new(Validator),
		Value:     value,
	}
}

func (v *FloatValidator) Max(maxValue float64) *FloatValidator {
	if v.Error != nil {
		return v
	}
	if v.Value > maxValue {
		v.Error = fmt.Errorf("%w: maximum value is %f", ErrMaxValueCheckFailed, maxValue)
	}
	return v
}

func (v *FloatValidator) Min(minValue float64) *FloatValidator {
	if v.Error != nil {
		return v
	}
	if v.Value < minValue {
		v.Error = fmt.Errorf("%w: minimum value is %f", ErrMinValueCheckFailed, minValue)
	}
	return v
}
