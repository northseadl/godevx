package validatex

import (
	"fmt"
	"regexp"
	"strings"
)

type StringValidator struct {
	*Validator
	value     string
	fieldName string
}

func NewStringValidator(value string) *StringValidator {
	return &StringValidator{
		Validator: new(Validator),
		value:     value,
	}
}

// Field 用于给验证器指定当前验证的字段名称。当产生验证错误时，该字段名会包含在错误信息中。
func (v *StringValidator) Field(name string) *StringValidator {
	v.fieldName = name
	return v
}

// 检查当前是否已有错误，如果有则直接返回当前对象以终止链式调用
func (v *StringValidator) checkError() bool {
	return v.Error != nil
}

// 通用的错误设置方法，允许包含字段名和格式化信息
func (v *StringValidator) fail(errType error, format string, args ...interface{}) *StringValidator {
	if v.Error == nil {
		prefix := ""
		if v.fieldName != "" {
			prefix = fmt.Sprintf("field '%s' ", v.fieldName)
		}
		v.Error = fmt.Errorf("%w: %s%s", errType, prefix, fmt.Sprintf(format, args...))
	}
	return v
}

func (v *StringValidator) MaxLen(n int) *StringValidator {
	if v.checkError() {
		return v
	}
	if len(v.value) > n {
		return v.fail(ErrMaxLenCheckFailed, "maximum length is %d, got length %d", n, len(v.value))
	}
	return v
}

func (v *StringValidator) MinLen(n int) *StringValidator {
	if v.checkError() {
		return v
	}
	if len(v.value) < n {
		return v.fail(ErrMinLenCheckFailed, "minimum length is %d, got length %d", n, len(v.value))
	}
	return v
}

func (v *StringValidator) Len(n int) *StringValidator {
	if v.checkError() {
		return v
	}
	if len(v.value) != n {
		return v.fail(ErrLenCheckFailed, "length must be %d, got length %d", n, len(v.value))
	}
	return v
}

func (v *StringValidator) Contains(str string) *StringValidator {
	if v.checkError() {
		return v
	}
	if !strings.Contains(v.value, str) {
		return v.fail(ErrStringCheckFailed, "value '%s' does not contain '%s'", v.value, str)
	}
	return v
}

func (v *StringValidator) HasPrefix(prefix string) *StringValidator {
	if v.checkError() {
		return v
	}
	if !strings.HasPrefix(v.value, prefix) {
		return v.fail(ErrPrefixCheckFailed, "value '%s' does not have prefix '%s'", v.value, prefix)
	}
	return v
}

func (v *StringValidator) HasSuffix(suffix string) *StringValidator {
	if v.checkError() {
		return v
	}
	if !strings.HasSuffix(v.value, suffix) {
		return v.fail(ErrSuffixCheckFailed, "value '%s' does not have suffix '%s'", v.value, suffix)
	}
	return v
}

func (v *StringValidator) MatchesRegex(regex string) *StringValidator {
	if v.checkError() {
		return v
	}
	match, err := regexp.MatchString(regex, v.value)
	if err != nil {
		return v.fail(ErrRegexCheckFailed, "regex '%s' matching failed: %v", regex, err)
	}
	if !match {
		return v.fail(ErrRegexCheckFailed, "value '%s' does not match regex '%s'", v.value, regex)
	}
	return v
}

func (v *StringValidator) IsEmail() *StringValidator {
	if v.checkError() {
		return v
	}
	if !regexEmail.MatchString(v.value) {
		return v.fail(ErrRegexCheckFailed, "value '%s' is not a valid email", v.value)
	}
	return v
}

func (v *StringValidator) IsURL() *StringValidator {
	if v.checkError() {
		return v
	}
	if !regexURL.MatchString(v.value) {
		return v.fail(ErrRegexCheckFailed, "value '%s' is not a valid URL", v.value)
	}
	return v
}

func (v *StringValidator) IsPhone() *StringValidator {
	if v.checkError() {
		return v
	}
	if !regexPhoneNumber.MatchString(v.value) {
		return v.fail(ErrRegexCheckFailed, "value '%s' is not a valid phone number", v.value)
	}
	return v
}

func (v *StringValidator) IsIP() *StringValidator {
	if v.checkError() {
		return v
	}
	if !regexIP.MatchString(v.value) {
		return v.fail(ErrRegexCheckFailed, "value '%s' is not a valid IP address", v.value)
	}
	return v
}
