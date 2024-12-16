package validatex

import (
	"fmt"
)

type ArrayValidator struct {
	*Validator
	value     []any
	fieldName string
}

func NewArrayValidator(value []any) *ArrayValidator {
	return &ArrayValidator{
		Validator: new(Validator),
		value:     value,
	}
}

// Field 为验证器指定当前验证的字段名称。当产生验证错误时，该字段名会包含在错误信息中。
func (v *ArrayValidator) Field(name string) *ArrayValidator {
	v.fieldName = name
	return v
}

// 检查当前是否已有错误，如果有则返回 true，以便中断后续检查
func (v *ArrayValidator) checkError() bool {
	return v.Error != nil
}

// 通用错误处理方法，包含字段名信息
func (v *ArrayValidator) fail(errType error, format string, args ...interface{}) *ArrayValidator {
	if v.Error == nil {
		prefix := ""
		if v.fieldName != "" {
			prefix = fmt.Sprintf("field '%s' ", v.fieldName)
		}
		v.Error = fmt.Errorf("%w: %s%s", errType, prefix, fmt.Sprintf(format, args...))
	}
	return v
}

func (v *ArrayValidator) MaxLen(n int) *ArrayValidator {
	if v.checkError() {
		return v
	}
	if len(v.value) > n {
		return v.fail(ErrMaxLenCheckFailed, "maximum length is %d, got length %d", n, len(v.value))
	}
	return v
}

func (v *ArrayValidator) MinLen(n int) *ArrayValidator {
	if v.checkError() {
		return v
	}
	if len(v.value) < n {
		return v.fail(ErrMinLenCheckFailed, "minimum length is %d, got length %d", n, len(v.value))
	}
	return v
}

// Items 接受一个函数对每个 array 元素进行验证。
// fn 回调函数接收当前元素和验证器本身。用户可在 fn 内调用 v.Error 设置错误。
func (v *ArrayValidator) Items(fn func(item any, validator *Validator)) *ArrayValidator {
	if v.checkError() {
		return v
	}

	for _, item := range v.value {
		fn(item, v.Validator)
		if v.checkError() {
			// 如果需要在出现错误后立即终止检查，可以在这里选择 break
			// break
			return v
		}
	}
	return v
}
