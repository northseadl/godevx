package debugx

import (
	"encoding/json"
	"fmt"
)

const debugOutputPrefix = "[DEBUG] "

func debugPrintln(format string, values ...any) {
	fmt.Printf(debugOutputPrefix+format+"\r\n", values...)
}

func debugPrintf(format string, values ...any) {
	fmt.Printf(debugOutputPrefix+format, values...)
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
func Println(values ...interface{}) {
	debugPrintln("%v", values...)
}

// Printf prints the given value with format.
//
//	将给定的值按格式打印出来。
func Printf(format string, values ...interface{}) {
	debugPrintf(format, values...)
}
