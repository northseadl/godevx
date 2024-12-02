package jsonx

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrInvalidType = errors.New("invalid type conversion")
	ErrKeyNotFound = errors.New("key not found in map")
)

// ParseJsonToMap 解析JSON字符串到map[string]any
func ParseJsonToMap(jsonStr string) (map[string]any, error) {
	if jsonStr == "" {
		return nil, ErrEmptyValue
	}

	var result map[string]any
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return nil, err
	}
	return result, nil
}

// ParseJsonToSlice 解析JSON字符串到[]any
func ParseJsonToSlice(jsonStr string) ([]any, error) {
	if jsonStr == "" {
		return nil, ErrEmptyValue
	}

	var result []any
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return nil, err
	}
	return result, nil
}

// JSONMap 提供便捷的map操作方法
type JSONMap map[string]any

// ParseToJSONMap 解析JSON字符串到JSONMap
func ParseToJSONMap(jsonStr string) (JSONMap, error) {
	m, err := ParseJsonToMap(jsonStr)
	if err != nil {
		return nil, err
	}
	return JSONMap(m), nil
}

// GetString 获取字符串值
func (m JSONMap) GetString(key string) (string, error) {
	v, ok := m[key]
	if !ok {
		return "", ErrKeyNotFound
	}

	str, ok := v.(string)
	if !ok {
		return "", ErrInvalidType
	}
	return str, nil
}

// GetInt 获取整数值
func (m JSONMap) GetInt(key string) (int64, error) {
	v, ok := m[key]
	if !ok {
		return 0, ErrKeyNotFound
	}

	switch v := v.(type) {
	case float64:
		return int64(v), nil
	case int64:
		return v, nil
	case string:
		return strconv.ParseInt(v, 10, 64)
	default:
		return 0, ErrInvalidType
	}
}

// GetFloat 获取浮点值
func (m JSONMap) GetFloat(key string) (float64, error) {
	v, ok := m[key]
	if !ok {
		return 0, ErrKeyNotFound
	}

	switch v := v.(type) {
	case float64:
		return v, nil
	case string:
		return strconv.ParseFloat(v, 64)
	default:
		return 0, ErrInvalidType
	}
}

// GetBool 获取布尔值
func (m JSONMap) GetBool(key string) (bool, error) {
	v, ok := m[key]
	if !ok {
		return false, ErrKeyNotFound
	}

	b, ok := v.(bool)
	if !ok {
		return false, ErrInvalidType
	}
	return b, nil
}

// GetMap 获取嵌套的map
func (m JSONMap) GetMap(key string) (JSONMap, error) {
	v, ok := m[key]
	if !ok {
		return nil, ErrKeyNotFound
	}

	nestedMap, ok := v.(map[string]any)
	if !ok {
		return nil, ErrInvalidType
	}
	return JSONMap(nestedMap), nil
}

// GetSlice 获取切片
func (m JSONMap) GetSlice(key string) ([]any, error) {
	v, ok := m[key]
	if !ok {
		return nil, ErrKeyNotFound
	}

	slice, ok := v.([]any)
	if !ok {
		return nil, ErrInvalidType
	}
	return slice, nil
}

// ParseWithValidator 使用验证器解析JSON
func ParseWithValidator(jsonStr string, v any) error {
	if err := json.Unmarshal([]byte(jsonStr), v); err != nil {
		return err
	}

	// 如果实现了Validator接口，则进行验证
	if validator, ok := v.(interface{ Validate() error }); ok {
		return validator.Validate()
	}
	return nil
}

// GetValueByPath 通过点分隔的路径获取嵌套值
func (m JSONMap) GetValueByPath(path string) (any, error) {
	keys := strings.Split(path, ".")
	current := m

	for i, key := range keys[:len(keys)-1] {
		nextMap, err := current.GetMap(key)
		if err != nil {
			return nil, fmt.Errorf("error at path %s: %w", strings.Join(keys[:i+1], "."), err)
		}
		current = nextMap
	}

	lastKey := keys[len(keys)-1]
	value, exists := current[lastKey]
	if !exists {
		return nil, ErrKeyNotFound
	}
	return value, nil
}

// MustParseJsonToMap 解析JSON字符串到map，失败时panic
func MustParseJsonToMap(jsonStr string) map[string]any {
	result, err := ParseJsonToMap(jsonStr)
	if err != nil {
		panic(err)
	}
	return result
}
