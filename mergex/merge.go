package mergex

import (
	"errors"
	"reflect"
)

var (
	ErrInvalidInput = errors.New("invalid input: source or destination must be struct")
	ErrNilPointer   = errors.New("destination cannot be nil pointer")
)

// MergeOption 定义合并选项
type MergeOption struct {
	// 是否忽略零值
	IgnoreZero bool
	// 是否深度合并
	DeepMerge bool
	// 自定义字段映射
	FieldMapping map[string]string
	// 字段过滤函数
	FieldFilter func(field reflect.StructField) bool
}

// DefaultOption 返回默认选项
func DefaultOption() *MergeOption {
	return &MergeOption{
		IgnoreZero:   false,
		DeepMerge:    false,
		FieldMapping: make(map[string]string),
	}
}

// MergeTo 增强版合并函数
func MergeTo[S any, D any](src S, dest *D, opts ...*MergeOption) error {
	if dest == nil {
		return ErrNilPointer
	}

	opt := DefaultOption()
	if len(opts) > 0 {
		opt = opts[0]
	}

	s := reflect.ValueOf(src)
	d := reflect.ValueOf(dest).Elem()

	if s.Kind() != reflect.Struct || d.Kind() != reflect.Struct {
		return ErrInvalidInput
	}

	return mergeFields(s, d, opt)
}

// mergeFields 合并字段实现
func mergeFields(src, dest reflect.Value, opt *MergeOption) error {
	for i := 0; i < src.NumField(); i++ {
		sField := src.Field(i)
		sType := src.Type().Field(i)

		// 应用字段过滤
		if opt.FieldFilter != nil && !opt.FieldFilter(sType) {
			continue
		}

		// 获取目标字段名
		destFieldName := sType.Name
		if mapped, ok := opt.FieldMapping[destFieldName]; ok {
			destFieldName = mapped
		}

		dField := dest.FieldByName(destFieldName)
		if !dField.IsValid() || !dField.CanSet() {
			continue
		}

		// 处理深度合并
		if opt.DeepMerge && sField.Kind() == reflect.Struct {
			mergeFields(sField, dField, opt)
			continue
		}

		// 处理值合并
		if shouldMergeField(dField, sField, opt) {
			dField.Set(sField)
		}
	}

	return nil
}

// shouldMergeField 判断是否应该合并字段
func shouldMergeField(dest, src reflect.Value, opt *MergeOption) bool {
	if !dest.IsValid() || dest.Kind() != src.Kind() {
		return false
	}

	if opt.IgnoreZero && src.IsZero() {
		return false
	}

	return true
}

// Clone 深度克隆结构体
func Clone[T any](src T) (T, error) {
	var dest T
	err := MergeTo(src, &dest, &MergeOption{DeepMerge: true})
	return dest, err
}
