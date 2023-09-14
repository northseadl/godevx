package validatex

import "errors"

var (
	ErrTypeInvalid       = errors.New("invalid type")
	ErrMaxLenCheckFailed = errors.New("maximum length validation failed")
	ErrMinLenCheckFailed = errors.New("minimum length validation failed")
	ErrStringCheckFailed = errors.New("string validation failed")
	ErrPrefixCheckFailed = errors.New("prefix validation failed")
	ErrSuffixCheckFailed = errors.New("suffix validation failed")
	ErrRegexCheckFailed  = errors.New("regex validation failed")
)

var (
	ErrMaxValueCheckFailed = errors.New("maximum value validation failed")
	ErrMinValueCheckFailed = errors.New("minimum value validation failed")
	ErrInCheckFailed       = errors.New("in validation failed")
	ErrNotInCheckFailed    = errors.New("not in validation failed")
)
