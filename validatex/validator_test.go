package validatex

import (
	"fmt"
	"testing"
)

func TestValidatorString(t *testing.T) {
	testCases := []struct {
		value         string
		maxLen        int
		minLen        int
		contains      string
		hasPrefix     string
		hasSuffix     string
		matchesRegex  string
		expectedError error
	}{
		{"hello", 10, 3, "e", "h", "o", "", nil},
		{"hello", 4, 3, "e", "h", "o", "", fmt.Errorf("%w: maximum length is %d", ErrMaxLenCheckFailed, 4)},
		{"hello", 10, 6, "e", "h", "o", "", fmt.Errorf("%w: minimum length is %d", ErrMinLenCheckFailed, 6)},
		{"hello", 10, 3, "x", "h", "o", "", fmt.Errorf("%w: does not contain '%s'", ErrStringCheckFailed, "x")},
		{"hello", 10, 3, "e", "x", "o", "", fmt.Errorf("%w: does not have prefix '%s'", ErrPrefixCheckFailed, "x")},
		{"hello", 10, 3, "e", "h", "x", "", fmt.Errorf("%w: does not have suffix '%s'", ErrSuffixCheckFailed, "x")},
		{"hello", 10, 3, "e", "h", "o", "^[a-z]+$", nil},
		{"hello", 10, 3, "e", "h", "o", "^[0-9]+$", fmt.Errorf("%w: regex match failed", ErrRegexCheckFailed)},
	}

	for _, tc := range testCases {
		v := NewValidator().String(tc.value)
		v.MaxLen(tc.maxLen).MinLen(tc.minLen).Contains(tc.contains).HasPrefix(tc.hasPrefix).HasSuffix(tc.hasSuffix).MatchesRegex(tc.matchesRegex)
		if tc.expectedError == nil && v.Error != nil {
			t.Errorf("Unexpected error: %v", v.Error)
		} else if tc.expectedError != nil && (v.Error == nil || v.Error.Error() != tc.expectedError.Error()) {
			t.Errorf("Expected error: %s, got: %v", tc.expectedError, v.Error)
		}
	}
}

func TestValidatorUInt(t *testing.T) {
	testCases := []struct {
		value         uint64
		min           uint64
		max           uint64
		in            []uint64
		notIn         []uint64
		expectedError error
	}{
		{5, 1, 10, []uint64{1, 2, 3, 4, 5}, []uint64{6, 7, 8, 9, 10}, nil},
		{5, 6, 10, []uint64{1, 2, 3, 4, 5}, []uint64{6, 7, 8, 9, 10}, fmt.Errorf("%w: minimum value is %d", ErrMinValueCheckFailed, 6)},
		{5, 1, 4, []uint64{1, 2, 3, 4, 5}, []uint64{6, 7, 8, 9, 10}, fmt.Errorf("%w: maximum value is %d", ErrMaxValueCheckFailed, 4)},
		{5, 1, 10, []uint64{1, 2, 3, 4}, []uint64{6, 7, 8, 9, 10}, fmt.Errorf("%w: value must be in %v", ErrInCheckFailed, []uint64{1, 2, 3, 4})},
		{5, 1, 10, []uint64{1, 2, 3, 4, 5}, []uint64{5, 6, 7, 8, 9, 10}, ErrNotInCheckFailed},
	}

	for _, tc := range testCases {
		v := NewValidator().UInt(tc.value)
		v.Min(tc.min).Max(tc.max).In(tc.in...).NotIn(tc.notIn...)
		if tc.expectedError == nil && v.Error != nil {
			t.Errorf("Unexpected error: %v", v.Error)
		} else if tc.expectedError != nil && (v.Error == nil || v.Error.Error() != tc.expectedError.Error()) {
			t.Errorf("Expected error: %s, got: %v", tc.expectedError, v.Error)
		}
	}
}

func TestValidatorInt(t *testing.T) {
	testCases := []struct {
		value         int64
		min           int64
		max           int64
		in            []int64
		notIn         []int64
		expectedError error
	}{
		{5, 1, 10, []int64{1, 2, 3, 4, 5}, []int64{6, 7, 8, 9, 10}, nil},
		{5, 6, 10, []int64{1, 2, 3, 4, 5}, []int64{6, 7, 8, 9, 10}, fmt.Errorf("%w: minimum value is %d", ErrMinValueCheckFailed, 6)},
		{5, 1, 4, []int64{1, 2, 3, 4, 5}, []int64{6, 7, 8, 9, 10}, fmt.Errorf("%w: maximum value is %d", ErrMaxValueCheckFailed, 4)},
		{5, 1, 10, []int64{1, 2, 3, 4}, []int64{6, 7, 8, 9, 10}, fmt.Errorf("%w: value must be in %v", ErrInCheckFailed, []int64{1, 2, 3, 4})},
		{5, 1, 10, []int64{1, 2, 3, 4, 5}, []int64{5, 6, 7, 8, 9, 10}, ErrNotInCheckFailed},
	}

	for _, tc := range testCases {
		v := NewValidator().Int(tc.value)
		v.Min(tc.min).Max(tc.max).In(tc.in...).NotIn(tc.notIn...)
		if tc.expectedError == nil && v.Error != nil {
			t.Errorf("Unexpected error: %v", v.Error)
		} else if tc.expectedError != nil && (v.Error == nil || v.Error.Error() != tc.expectedError.Error()) {
			t.Errorf("Expected error: %s, got: %v", tc.expectedError, v.Error)
		}
	}
}

func TestValidatorFloat(t *testing.T) {
	testCases := []struct {
		value         float64
		min           float64
		max           float64
		expectedError error
	}{
		{5.5, 1.0, 10.0, nil},
		{5.5, 6.0, 10.0, fmt.Errorf("%w: minimum value is %f", ErrMinValueCheckFailed, 6.0)},
		{5.5, 1.0, 4.0, fmt.Errorf("%w: maximum value is %f", ErrMaxValueCheckFailed, 4.0)},
	}

	for _, tc := range testCases {
		v := NewValidator().Float(tc.value)
		v.Min(tc.min).Max(tc.max)
		if tc.expectedError == nil && v.Error != nil {
			t.Errorf("Unexpected error: %v", v.Error)
		} else if tc.expectedError != nil && (v.Error == nil || v.Error.Error() != tc.expectedError.Error()) {
			t.Errorf("Expected error: %s, got: %v", tc.expectedError, v.Error)
		}
	}
}

func TestValidatorArray(t *testing.T) {
	testCases := []struct {
		value         []any
		minLen        int
		maxLen        int
		expectedError error
	}{
		{[]any{1, 2, 3}, 1, 5, nil},
		{[]any{1, 2, 3}, 4, 5, fmt.Errorf("%w: minimum length is %d", ErrMinLenCheckFailed, 4)},
		{[]any{1, 2, 3}, 1, 2, fmt.Errorf("%w: maximum length is %d", ErrMaxLenCheckFailed, 2)},
	}

	for _, tc := range testCases {
		v := NewValidator().Array(tc.value)
		v.MinLen(tc.minLen).MaxLen(tc.maxLen)
		if tc.expectedError == nil && v.Error != nil {
			t.Errorf("Unexpected error: %v", v.Error)
		} else if tc.expectedError != nil && (v.Error == nil || v.Error.Error() != tc.expectedError.Error()) {
			t.Errorf("Expected error: %s, got: %v", tc.expectedError, v.Error)
		}
	}
}
