package debugx

import (
	"encoding/json"
	"fmt"
)

const debugOutputPrefix = "[DEBUG] "

func debugPrintln(format string, values ...any) {
	fmt.Printf(debugOutputPrefix+format+"\r\n", values...)
}

func PrintJson(value any) {
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		debugPrintln("%v", value)
	}
	debugPrintln("%s", string(jsonBytes))
}

func Print() {

}
