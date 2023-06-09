package jsonx

import "encoding/json"

// MustMarshalToString marshals the given value to a string.
//
//	将给定的值序列化为一个字符串。
func MustMarshalToString(value any) string {
	bytes, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

// MustMarshalToBytes marshals the given value to a byte slice.
//
//	将给定的值序列化为一个字节切片。
func MustMarshalToBytes(value any) []byte {
	bytes, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	return bytes
}
