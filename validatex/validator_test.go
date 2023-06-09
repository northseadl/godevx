package validatex

import (
	"testing"
)

var caseString = []struct {
	value         string
	maxLen        int
	minLen        int
	contains      string
	prefix        string
	expectedError error
}{
	{"test", 8, 2, "est", "te", nil},
}

func TestValidator_String(t *testing.T) {
	for _, c := range caseString {
		v := NewValidator()
		v.String(c.value).MaxLen(c.maxLen).MinLen(c.minLen).Contains(c.contains).HasPrefix(c.prefix)
		if v.Error != c.expectedError {
			t.Errorf("Expected error %v, got %v", c.expectedError, v.Error)
		}
	}
}
