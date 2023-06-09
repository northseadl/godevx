package validatex

import (
	"github.com/northseadl/godevx/validatex/validators"
)

type Validator struct {
	Error error
}

func NewValidator() *Validator {
	return new(Validator)
}

func (v *Validator) String(value string) *validators.StringValidator {
	return &validators.StringValidator{
		Validator: v,
		Value:     value,
	}
}

func (v *Validator) Int(value any) *validators.IntValidator {
	var wrapValue int64
	switch tv := value.(type) {
	case int8:
		wrapValue = int64(tv)
	case int16:
		wrapValue = int64(tv)
	case int32:
		wrapValue = int64(tv)
	case int:
		wrapValue = int64(tv)
	case int64:
		wrapValue = tv
	default:
		v.Error = validators.ErrTypeInvalid
	}
	return &validators.IntValidator{
		Validator: v,
		Value:     wrapValue,
	}
}

func (v *Validator) UInt(value any) *validators.UIntValidator {
	var wrapValue uint64
	switch tv := value.(type) {
	case uint8:
		wrapValue = uint64(tv)
	case uint16:
		wrapValue = uint64(tv)
	case uint32:
		wrapValue = uint64(tv)
	case uint:
		wrapValue = uint64(tv)
	case uint64:
		wrapValue = tv
	default:
		v.Error = validators.ErrTypeInvalid
	}
	return &validators.UIntValidator{
		Validator: v,
		Value:     wrapValue,
	}
}

func (v *Validator) Float(value any) *validators.FloatValidator {
	var wrapValue float64
	switch tv := value.(type) {
	case float32:
		wrapValue = float64(tv)
	case float64:
		wrapValue = tv
	default:
		v.Error = validators.ErrTypeInvalid
	}
	return &validators.FloatValidator{
		Validator: v,
		Value:     wrapValue,
	}
}
