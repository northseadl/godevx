package validators

import (
	"fmt"
	"github.com/northseadl/godevx/validatex"
)

type IntValidator struct {
	*validatex.Validator
	Value int64
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
