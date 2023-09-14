package validatex

import (
	"fmt"
)

type IntValidator struct {
	*Validator
	Value int64
}

func NewIntValidator(value int64) *IntValidator {
	return &IntValidator{
		Validator: new(Validator),
		Value:     value,
	}
}

func (v *IntValidator) Max(maxValue int64) *IntValidator {
	if v.Error != nil {
		return v
	}
	if v.Value > maxValue {
		v.Error = fmt.Errorf("%w: maximum value is %d", ErrMaxValueCheckFailed, maxValue)
	}
	return v
}

func (v *IntValidator) Min(minValue int64) *IntValidator {
	if v.Error != nil {
		return v
	}
	if v.Value < minValue {
		v.Error = fmt.Errorf("%w: minimum value is %d", ErrMinValueCheckFailed, minValue)
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
	v.Error = fmt.Errorf("%w: value must be in %v", ErrInCheckFailed, values)
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
