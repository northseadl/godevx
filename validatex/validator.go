package validatex

type Validator struct {
	Error error
}

func NewValidator() *Validator {
	return new(Validator)
}

func (v *Validator) String(value string) *StringValidator {
	return &StringValidator{
		Validator: v,
		value:     value,
	}
}

func (v *Validator) Int(value any) *IntValidator {
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
		v.Error = ErrTypeInvalid
	}
	return &IntValidator{
		Validator: v,
		Value:     wrapValue,
	}
}

func (v *Validator) UInt(value any) *UIntValidator {
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
		v.Error = ErrTypeInvalid
	}
	return &UIntValidator{
		Validator: v,
		value:     wrapValue,
	}
}

func (v *Validator) Float(value any) *FloatValidator {
	var wrapValue float64
	switch tv := value.(type) {
	case float32:
		wrapValue = float64(tv)
	case float64:
		wrapValue = tv
	default:
		v.Error = ErrTypeInvalid
	}
	return &FloatValidator{
		Validator: v,
		Value:     wrapValue,
	}
}

func (v *Validator) Array(values []any) *ArrayValidator {
	return &ArrayValidator{
		Validator: v,
		value:     values,
	}
}
