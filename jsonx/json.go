package jsonx

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"os"
)

// JSON编码相关错误
var (
	ErrInvalidJSON = errors.New("invalid json format")
	ErrEmptyValue  = errors.New("empty value")
)

// Marshal 将值序列化为JSON字节切片,支持错误处理
func Marshal(v any) ([]byte, error) {
	if v == nil {
		return nil, ErrEmptyValue
	}
	return json.Marshal(v)
}

// MustMarshalToString 将值序列化为JSON字符串,遇错误则panic
func MustMarshalToString(v any) string {
	bytes, err := Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

// MustMarshalToBytes 将值序列化为JSON字节切片,遇错误则panic
func MustMarshalToBytes(v any) []byte {
	bytes, err := Marshal(v)
	if err != nil {
		panic(err)
	}
	return bytes
}

// Unmarshal 将JSON数据解析到指定结构体
func Unmarshal(data []byte, v any) error {
	if len(data) == 0 {
		return ErrEmptyValue
	}
	return json.Unmarshal(data, v)
}

// MustUnmarshal 将JSON数据解析到指定结构体,遇错误则panic
func MustUnmarshal(data []byte, v any) {
	if err := Unmarshal(data, v); err != nil {
		panic(err)
	}
}

// PrettyMarshal 将值序列化为格式化的JSON字符串
func PrettyMarshal(v any) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}

// WriteToFile 将值序列化并写入文件
func WriteToFile(filename string, v any) error {
	data, err := Marshal(v)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// ReadFromFile 从文件读取JSON并解析到结构体
func ReadFromFile(filename string, v any) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return Unmarshal(data, v)
}

// ValidateJSON 验证JSON字符串是否合法
func ValidateJSON(data []byte) bool {
	return json.Valid(data)
}

// Decode 从io.Reader解码JSON数据
func Decode(r io.Reader, v any) error {
	return json.NewDecoder(r).Decode(v)
}

// Encode 将数据编码为JSON并写入io.Writer
func Encode(w io.Writer, v any) error {
	return json.NewEncoder(w).Encode(v)
}

// DeepCopy 通过JSON序列化实现深拷贝
func DeepCopy(src, dst any) error {
	data, err := Marshal(src)
	if err != nil {
		return err
	}
	return Unmarshal(data, dst)
}

// CompactJSON 压缩JSON字符串,删除空白字符
func CompactJSON(data []byte) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := json.Compact(buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
