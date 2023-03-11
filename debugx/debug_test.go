package debugx

import "testing"

var case1 = struct {
	A string
	B int
	C []string
}{
	A: "case1",
	B: 1,
	C: []string{"case1", "debug"},
}

func TestPrintJson(t *testing.T) {
	PrintJson(case1)
	// todo test
}
