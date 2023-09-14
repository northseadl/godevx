package validatex

import (
	"fmt"
	"regexp"
	"strings"
)

type StringValidator struct {
	*Validator
	value string
}

func NewStringValidator(value string) *StringValidator {
	return &StringValidator{
		Validator: new(Validator),
		value:     value,
	}
}

func (v *StringValidator) MaxLen(n int) *StringValidator {
	if v.Error != nil {
		return v
	}
	if len(v.value) > n {
		v.Error = fmt.Errorf("%w: maximum length is %d", ErrMaxLenCheckFailed, n)
	}
	return v
}

func (v *StringValidator) MinLen(n int) *StringValidator {
	if v.Error != nil {
		return v
	}
	if len(v.value) < n {
		v.Error = fmt.Errorf("%w: minimum length is %d", ErrMinLenCheckFailed, n)
	}
	return v
}

func (v *StringValidator) Contains(str string) *StringValidator {
	if v.Error != nil {
		return v
	}
	if !strings.Contains(v.value, str) {
		v.Error = fmt.Errorf("%w: does not contain '%s'", ErrStringCheckFailed, str)
	}
	return v
}

func (v *StringValidator) HasPrefix(prefix string) *StringValidator {
	if v.Error != nil {
		return v
	}
	if !strings.HasPrefix(v.value, prefix) {
		v.Error = fmt.Errorf("%w: does not have prefix '%s'", ErrPrefixCheckFailed, prefix)
	}
	return v
}

func (v *StringValidator) HasSuffix(suffix string) *StringValidator {
	if v.Error != nil {
		return v
	}
	if !strings.HasSuffix(v.value, suffix) {
		v.Error = fmt.Errorf("%w: does not have suffix '%s'", ErrSuffixCheckFailed, suffix)
	}
	return v
}

func (v *StringValidator) MatchesRegex(regex string) *StringValidator {
	if v.Error != nil {
		return v
	}
	match, err := regexp.MatchString(regex, v.value)
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
	if !regexEmail.MatchString(v.value) {
		v.Error = fmt.Errorf("%w: regex match email failed", ErrRegexCheckFailed)
	}
	return v
}

func (v *StringValidator) IsURL() *StringValidator {
	if v.Error != nil {
		return v
	}
	if !regexURL.MatchString(v.value) {
		v.Error = fmt.Errorf("%w: regex match url failed", ErrRegexCheckFailed)
	}
	return v
}

func (v *StringValidator) IsPhone() *StringValidator {
	if v.Error != nil {
		return v
	}
	if !regexPhoneNumber.MatchString(v.value) {
		v.Error = fmt.Errorf("%w: regex match phone failed", ErrRegexCheckFailed)
	}
	return v
}

func (v *StringValidator) IsIP() *StringValidator {
	if v.Error != nil {
		return v
	}
	if !regexIP.MatchString(v.value) {
		v.Error = fmt.Errorf("%w: regex match ip failed", ErrRegexCheckFailed)
	}
	return v
}
