package validatex

import "godevx/validatex/validators"

type Validator struct {
	Error error
}

func (v *Validator) String(value string) *validators.StringValidator {
	return &validators.StringValidator{
		Validator: v,
		Value:     value,
	}
}
