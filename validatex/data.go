package validatex

type Data interface {
	Validate(func(validator Validator) error) error
}
