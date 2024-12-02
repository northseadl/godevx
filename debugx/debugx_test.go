// debugx/debugx_test.go
package debugx

import (
	"bytes"
	"strings"
	"testing"
)

func TestPrintln(t *testing.T) {
	// 准备测试buffer
	buf := &bytes.Buffer{}
	SetWriter(buf)

	tests := []struct {
		name     string
		input    []interface{}
		expected string
	}{
		{
			name:     "print string",
			input:    []interface{}{"hello"},
			expected: debugOutputPrefix + "hello\n",
		},
		{
			name:     "print multiple values",
			input:    []interface{}{1, "world", true},
			expected: debugOutputPrefix + "1 world true\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf.Reset()
			Println(tt.input...)
			if got := buf.String(); got != tt.expected {
				t.Errorf("Println() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPrintJson(t *testing.T) {
	buf := &bytes.Buffer{}
	SetWriter(buf)

	type testStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	test := testStruct{
		Name: "test",
		Age:  18,
	}

	if err := PrintJson(test); err != nil {
		t.Errorf("PrintJson() error = %v", err)
	}

	expected := `{"name":"test","age":18}`
	if !strings.Contains(buf.String(), expected) {
		t.Errorf("PrintJson() = %v, want %v", buf.String(), expected)
	}
}

func TestEnable(t *testing.T) {
	buf := &bytes.Buffer{}
	SetWriter(buf)

	// 禁用调试输出
	Enable(false)
	Println("this should not be printed")
	if buf.Len() > 0 {
		t.Error("Expected no output when debug is disabled")
	}

	// 启用调试输出
	Enable(true)
	Println("this should be printed")
	if buf.Len() == 0 {
		t.Error("Expected output when debug is enabled")
	}
}
