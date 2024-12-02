package debugx

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
)

const (
	debugOutputPrefix = "[DEBUG] "
)

type Debugger struct {
	mu     sync.Mutex
	enable bool
	writer io.Writer
}

var defaultDebugger = &Debugger{
	enable: true,
	writer: os.Stdout,
}

// Enable 启用或禁用调试输出
func Enable(enable bool) {
	defaultDebugger.mu.Lock()
	defer defaultDebugger.mu.Unlock()
	defaultDebugger.enable = enable
}

// SetWriter 设置输出writer
func SetWriter(w io.Writer) {
	defaultDebugger.mu.Lock()
	defer defaultDebugger.mu.Unlock()
	defaultDebugger.writer = w
}

func (d *Debugger) println(values ...interface{}) {
	if !d.enable {
		return
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	_, _ = fmt.Fprint(d.writer, debugOutputPrefix)
	_, _ = fmt.Fprintln(d.writer, values...)
}

func (d *Debugger) printf(format string, values ...interface{}) {
	if !d.enable {
		return
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	_, _ = fmt.Fprintf(d.writer, debugOutputPrefix+format, values...)
}

// PrintJson 将给定的值打印为JSON
func PrintJson(value any) error {
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		defaultDebugger.println("marshal json error:", err, "value:", value)
		return err
	}
	defaultDebugger.println(string(jsonBytes))
	return nil
}

// PrintJsonPretty 将给定的值打印为带缩进的JSON
func PrintJsonPretty(value any) error {
	jsonBytes, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		defaultDebugger.println("marshal json error:", err, "value:", value)
		return err
	}
	defaultDebugger.println(string(jsonBytes))
	return nil
}

// Println 将给定的值打印出来
func Println(values ...interface{}) {
	defaultDebugger.println(values...)
}

// Printf 将给定的值按格式打印出来
func Printf(format string, values ...interface{}) {
	defaultDebugger.printf(format, values...)
}
