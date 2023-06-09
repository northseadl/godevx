package validators

import (
	"fmt"
	"github.com/northseadl/godevx/validatex"
)

type UIntValidator struct {
	Validator *validatex.Validator
	Value     uint64
}

func NewUIntValidator(value uint64) *UIntValidator {
	return &UIntValidator{
		Validator: new(validatex.Validator),
		Value:     value,
	}
}

// Min checks if the Value is greater than or equal to the given minimum Value.
func (v *UIntValidator) Min(min uint64) *UIntValidator {
	if v.Validator.Error != nil {
		return v
	}
	if v.Value < min {
		v.Validator.Error = fmt.Errorf("%w: minimum Value is %d", ErrMinValueCheckFailed, min)
	}
	return v
}

// Max checks if the Value is less than or equal to the given maximum Value.
func (v *UIntValidator) Max(max uint64) *UIntValidator {
	if v.Validator.Error != nil {
		return v
	}
	if v.Value > max {
		v.Validator.Error = fmt.Errorf("%w: maximum Value is %d", ErrMaxValueCheckFailed, max)
	}
	return v
}

// In checks if the Value is in the given list of values.
func (v *UIntValidator) In(values ...uint64) *UIntValidator {
	if v.Validator.Error != nil {
		return v
	}
	for _, value := range values {
		if v.Value == value {
			return v
		}
	}
	v.Validator.Error = fmt.Errorf("%w: Value must be in %v", ErrInCheckFailed, values)
	return v
}

// NotIn checks if the Value is not in the given list of values.
func (v *UIntValidator) NotIn(values ...uint64) *UIntValidator {
	if v.Validator.Error != nil {
		return v
	}
	for _, value := range values {
		if v.Value == value {
			v.Validator.Error = ErrNotInCheckFailed
			return v
		}
	}
	return v
}
