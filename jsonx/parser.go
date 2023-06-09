package jsonx

import "encoding/json"

// ParseJsonToMap parses the given JSON string to a map.
//
//	解析给定的 JSON 字符串到一个 map。
func ParseJsonToMap(jsonStr string) (map[string]any, error) {
	var result map[string]any
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return nil, err
	}
	return result, nil
}
