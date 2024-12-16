package validatex

import (
	"fmt"
)

type FloatValidator struct {
	*Validator
	Value     float64
	fieldName string
}

func NewFloatValidator(value float64) *FloatValidator {
	return &FloatValidator{
		Validator: new(Validator),
		Value:     value,
	}
}

// Field 用于给验证器指定当前验证的字段名称。当产生验证错误时，字段名会包含在错误信息中。
func (v *FloatValidator) Field(name string) *FloatValidator {
	v.fieldName = name
	return v
}

// 检查当前是否已有错误，如果有则直接返回当前对象以终止链式调用
func (v *FloatValidator) checkError() bool {
	return v.Error != nil
}

// 通用的错误设置方法，允许包含字段名和格式化信息
func (v *FloatValidator) fail(errType error, format string, args ...interface{}) *FloatValidator {
	if v.Error == nil {
		prefix := ""
		if v.fieldName != "" {
			prefix = fmt.Sprintf("field '%s' ", v.fieldName)
		}
		v.Error = fmt.Errorf("%w: %s%s", errType, prefix, fmt.Sprintf(format, args...))
	}
	return v
}

func (v *FloatValidator) Max(maxValue float64) *FloatValidator {
	if v.checkError() {
		return v
	}
	if v.Value > maxValue {
		return v.fail(ErrMaxValueCheckFailed, "maximum value is %f, got %f", maxValue, v.Value)
	}
	return v
}

func (v *FloatValidator) Min(minValue float64) *FloatValidator {
	if v.checkError() {
		return v
	}
	if v.Value < minValue {
		return v.fail(ErrMinValueCheckFailed, "minimum value is %f, got %f", minValue, v.Value)
	}
	return v
}
