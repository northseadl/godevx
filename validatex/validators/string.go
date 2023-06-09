package validators

import (
	"fmt"
	"github.com/northseadl/godevx/validatex"
	"regexp"
	"strings"
)

type StringValidator struct {
	*validatex.Validator
	Value string
}

func NewStringValidator(value string) *StringValidator {
	return &StringValidator{
		Validator: new(validatex.Validator),
		Value:     value,
	}
}

func (v *StringValidator) MaxLen(n int) *StringValidator {
	if v.Error != nil {
		return v
	}
	if len(v.Value) > n {
		v.Error = fmt.Errorf("%w: maximum length is %d", ErrMaxLenCheckFailed, n)
	}
	return v
}

func (v *StringValidator) MinLen(n int) *StringValidator {
	if v.Error != nil {
		return v
	}
	if len(v.Value) < n {
		v.Error = fmt.Errorf("%w: minimum length is %d", ErrMinLenCheckFailed, n)
	}
	return v
}

func (v *StringValidator) Contains(str string) *StringValidator {
	if v.Error != nil {
		return v
	}
	if !strings.Contains(v.Value, str) {
		v.Error = fmt.Errorf("%w: does not contain '%s'", ErrStringCheckFailed, str)
	}
	return v
}

func (v *StringValidator) HasPrefix(prefix string) *StringValidator {
	if v.Error != nil {
		return v
	}
	if !strings.HasPrefix(v.Value, prefix) {
		v.Error = fmt.Errorf("%w: does not have prefix '%s'", ErrPrefixCheckFailed, prefix)
	}
	return v
}

func (v *StringValidator) HasSuffix(suffix string) *StringValidator {
	if v.Error != nil {
		return v
	}
	if !strings.HasSuffix(v.Value, suffix) {
		v.Error = fmt.Errorf("%w: does not have suffix '%s'", ErrSuffixCheckFailed, suffix)
	}
	return v
}

func (v *StringValidator) MatchesRegex(regex string) *StringValidator {
	if v.Error != nil {
		return v
	}
	match, err := regexp.MatchString(regex, v.Value)
	if err != nil {
		v.Error = fmt.Errorf("%w: regex matching failed", ErrRegexCheckFailed)
		return v
	}
	if !match {
		v.Error = fmt.Errorf("%w: regex match failed", ErrRegexCheckFailed)
	}
	return v
}

func (v *StringValidator) IsEmail() *StringValidator {
	if v.Error != nil {
		return v
	}
	if !regexEmail.MatchString(v.Value) {
		v.Error = fmt.Errorf("%w: regex match email failed", ErrRegexCheckFailed)
	}
	return v
}

func (v *StringValidator) IsURL() *StringValidator {
	if v.Error != nil {
		return v
	}
	if !regexURL.MatchString(v.Value) {
		v.Error = fmt.Errorf("%w: regex match url failed", ErrRegexCheckFailed)
	}
	return v
}

func (v *StringValidator) IsPhone() *StringValidator {
	if v.Error != nil {
		return v
	}
	if !regexPhoneNumber.MatchString(v.Value) {
		v.Error = fmt.Errorf("%w: regex match phone failed", ErrRegexCheckFailed)
	}
	return v
}

func (v *StringValidator) IsIP() *StringValidator {
	if v.Error != nil {
		return v
	}
	if !regexIP.MatchString(v.Value) {
		v.Error = fmt.Errorf("%w: regex match ip failed", ErrRegexCheckFailed)
	}
	return v
}
