package validatex

import (
	"fmt"
)

type UIntValidator struct {
	*Validator
	value uint64
}

func NewUIntValidator(value uint64) *UIntValidator {
	return &UIntValidator{
		Validator: new(Validator),
		value:     value,
	}
}

// Min checks if the value is greater than or equal to the given minimum value.
func (v *UIntValidator) Min(min uint64) *UIntValidator {
	if v.Validator.Error != nil {
		return v
	}
	if v.value < min {
		v.Validator.Error = fmt.Errorf("%w: minimum value is %d", ErrMinValueCheckFailed, min)
	}
	return v
}

// Max checks if the value is less than or equal to the given maximum value.
func (v *UIntValidator) Max(max uint64) *UIntValidator {
	if v.Validator.Error != nil {
		return v
	}
	if v.value > max {
		v.Validator.Error = fmt.Errorf("%w: maximum value is %d", ErrMaxValueCheckFailed, max)
	}
	return v
}

// In checks if the value is in the given list of values.
func (v *UIntValidator) In(values ...uint64) *UIntValidator {
	if v.Validator.Error != nil {
		return v
	}
	for _, value := range values {
		if v.value == value {
			return v
		}
	}
	v.Validator.Error = fmt.Errorf("%w: value must be in %v", ErrInCheckFailed, values)
	return v
}

// NotIn checks if the value is not in the given list of values.
func (v *UIntValidator) NotIn(values ...uint64) *UIntValidator {
	if v.Validator.Error != nil {
		return v
	}
	for _, value := range values {
		if v.value == value {
			v.Validator.Error = ErrNotInCheckFailed
			return v
		}
	}
	return v
}
