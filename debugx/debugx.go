package debugx

import (
	"encoding/json"
	"fmt"
)

const debugOutputPrefix = "[DEBUG] "

func debugPrintln(format string, values ...any) {
	fmt.Printf(debugOutputPrefix+format+"\r\n", values...)
}

// PrintJson prints the given value as JSON.
//
//	将给定的值打印为 JSON。
func PrintJson(value any) {
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		debugPrintln("%v", value)
	}
	debugPrintln("%s", string(jsonBytes))
}

// PrintJsonPretty prints the given value as JSON with indentation.
//
//	将给定的值打印为带缩进的 JSON。
func PrintJsonPretty(value any) {
	jsonBytes, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		debugPrintln("%v", value)
	}
	debugPrintln("%s", string(jsonBytes))
}

// Println prints the given value.
//
//	将给定的值打印出来。
func Println(value any) {
	debugPrintln("%v", value)
}
