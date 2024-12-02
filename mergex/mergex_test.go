package mergex

import (
	"github.com/northseadl/godevx/debugx"
	"testing"
)

var (
	case1 = struct {
		A string
		B int
		C []string
	}{
		A: "case1",
		B: 1,
		C: []string{"case1", "debug"},
	}

	case2 = struct {
		A string
		B int
	}{}

	case3 = struct {
		A string
		B string
		C []string
		D bool
	}{
		A: "case3",
	}
)

func c1() struct {
	A string
	B int
	C []string
} {
	return case1
}

func c2() struct {
	A string
	B int
} {
	return case2
}

func c3() struct {
	A string
	B string
	C []string
	D bool
} {
	return case3
}

func TestOverWriteTo(t *testing.T) {
	c2 := c2()
	OverWriteTo(case1, &c2)
	debugx.PrintJson(c2)
	c3 := c3()
	OverWriteTo(case1, &c3)
	debugx.PrintJson(c3)
	// todo test
}

func TestMergeTo(t *testing.T) {
	c2 := c2()
	MergeTo(case1, &c2)
	debugx.PrintJson(c2)
	c3 := c3()
	MergeTo(case1, &c3)
	debugx.PrintJson(c3)
	// todo test
}
