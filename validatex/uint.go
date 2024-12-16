package validatex

import (
	"fmt"
)

type UIntValidator struct {
	*Validator
	value     uint64
	fieldName string
}

func NewUIntValidator(value uint64) *UIntValidator {
	return &UIntValidator{
		Validator: new(Validator),
		value:     value,
	}
}

// Field 用于为验证器指定字段名。当产生验证错误时，该字段名将包含在错误信息中。
func (v *UIntValidator) Field(name string) *UIntValidator {
	v.fieldName = name
	return v
}

// 检查当前是否已有错误，如果有则返回true，以中断后续检查
func (v *UIntValidator) checkError() bool {
	return v.Error != nil
}

// 通用的错误设置方法，包含字段名和格式化信息
func (v *UIntValidator) fail(errType error, format string, args ...interface{}) *UIntValidator {
	if v.Error == nil {
		prefix := ""
		if v.fieldName != "" {
			prefix = fmt.Sprintf("field '%s' ", v.fieldName)
		}
		v.Error = fmt.Errorf("%w: %s%s", errType, prefix, fmt.Sprintf(format, args...))
	}
	return v
}

// Min checks if the value is greater than or equal to the given minimum value.
func (v *UIntValidator) Min(min uint64) *UIntValidator {
	if v.checkError() {
		return v
	}
	if v.value < min {
		return v.fail(ErrMinValueCheckFailed, "minimum value is %d, got %d", min, v.value)
	}
	return v
}

// Max checks if the value is less than or equal to the given maximum value.
func (v *UIntValidator) Max(max uint64) *UIntValidator {
	if v.checkError() {
		return v
	}
	if v.value > max {
		return v.fail(ErrMaxValueCheckFailed, "maximum value is %d, got %d", max, v.value)
	}
	return v
}

// In checks if the value is in the given list of values.
func (v *UIntValidator) In(values ...uint64) *UIntValidator {
	if v.checkError() {
		return v
	}
	for _, value := range values {
		if v.value == value {
			return v
		}
	}
	return v.fail(ErrInCheckFailed, "value must be in %v, got %d", values, v.value)
}

// NotIn checks if the value is not in the given list of values.
func (v *UIntValidator) NotIn(values ...uint64) *UIntValidator {
	if v.checkError() {
		return v
	}
	for _, value := range values {
		if v.value == value {
			return v.fail(ErrNotInCheckFailed, "value %d should not be in %v", v.value, values)
		}
	}
	return v
}
