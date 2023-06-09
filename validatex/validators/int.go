package validators

import (
	"fmt"
	"github.com/northseadl/godevx/validatex"
)

type IntValidator struct {
	*validatex.Validator
	Value int64
}

func NewIntValidator(value int64) *IntValidator {
	return &IntValidator{
		Validator: new(validatex.Validator),
		Value:     value,
	}
}

func (v *IntValidator) Max(maxValue int64) *IntValidator {
	if v.Error != nil {
		return v
	}
	if v.Value > maxValue {
		v.Error = fmt.Errorf("%w: maximum Value is %d", ErrMaxValueCheckFailed, maxValue)
	}
	return v
}

func (v *IntValidator) Min(minValue int64) *IntValidator {
	if v.Error != nil {
		return v
	}
	if v.Value < minValue {
		v.Error = fmt.Errorf("%w: minimum Value is %d", ErrMinValueCheckFailed, minValue)
	}
	return v
}

func (v *IntValidator) In(values ...int64) *IntValidator {
	if v.Error != nil {
		return v
	}
	for _, value := range values {
		if v.Value == value {
			return v
		}
	}
	v.Error = fmt.Errorf("%w: Value must be in %v", ErrInCheckFailed, values)
	return v
}

func (v *IntValidator) NotIn(values ...int64) *IntValidator {
	if v.Error != nil {
		return v
	}
	for _, value := range values {
		if v.Value == value {
			v.Error = ErrNotInCheckFailed
			return v
		}
	}
	return v
}
