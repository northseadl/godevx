package validatex

import (
	"fmt"
)

type IntValidator struct {
	*Validator
	Value     int64
	fieldName string
}

func NewIntValidator(value int64) *IntValidator {
	return &IntValidator{
		Validator: new(Validator),
		Value:     value,
	}
}

// Field 用于给验证器指定当前验证的字段名称。当产生验证错误时，该字段名会包含在错误信息中。
func (v *IntValidator) Field(name string) *IntValidator {
	v.fieldName = name
	return v
}

// 检查当前是否已有错误，如果有则直接返回当前对象以终止链式调用
func (v *IntValidator) checkError() bool {
	return v.Error != nil
}

// 通用的错误设置方法，允许包含字段名和格式化信息
func (v *IntValidator) fail(errType error, format string, args ...interface{}) *IntValidator {
	if v.Error == nil {
		prefix := ""
		if v.fieldName != "" {
			prefix = fmt.Sprintf("field '%s' ", v.fieldName)
		}
		v.Error = fmt.Errorf("%w: %s%s", errType, prefix, fmt.Sprintf(format, args...))
	}
	return v
}

func (v *IntValidator) Max(maxValue int64) *IntValidator {
	if v.checkError() {
		return v
	}
	if v.Value > maxValue {
		return v.fail(ErrMaxValueCheckFailed, "maximum value is %d, got %d", maxValue, v.Value)
	}
	return v
}

func (v *IntValidator) Min(minValue int64) *IntValidator {
	if v.checkError() {
		return v
	}
	if v.Value < minValue {
		return v.fail(ErrMinValueCheckFailed, "minimum value is %d, got %d", minValue, v.Value)
	}
	return v
}

func (v *IntValidator) In(values ...int64) *IntValidator {
	if v.checkError() {
		return v
	}
	for _, value := range values {
		if v.Value == value {
			return v
		}
	}
	return v.fail(ErrInCheckFailed, "value must be in %v, got %d", values, v.Value)
}

func (v *IntValidator) NotIn(values ...int64) *IntValidator {
	if v.checkError() {
		return v
	}
	for _, value := range values {
		if v.Value == value {
			return v.fail(ErrNotInCheckFailed, "value %d should not be in %v", v.Value, values)
		}
	}
	return v
}
