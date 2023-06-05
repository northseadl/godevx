package validatex

import (
	"github.com/northseadl/godevx/validatex/validators"
)

type Validator struct {
	Error error
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
