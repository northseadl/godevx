package slicex

import (
	"testing"
)

var (
	intCases = []struct {
		case1       []int
		case2       []int
		containWant bool
	}{
		{
			case1:       []int{1, 2, 3},
			case2:       []int{1},
			containWant: true,
		},
		{
			case1:       []int{1, 2, 3},
			case2:       []int{1, 2, 3},
			containWant: true,
		},
		{
			case1:       []int{1, 2, 3},
			case2:       []int{4},
			containWant: false,
		},
		{
			case1:       []int{1, 2, 3},
			case2:       []int{1, 2, 3, 4},
			containWant: false,
		},
	}
	int64Cases = []struct {
		case1        []int64
		case2        []int64
		containsWant bool
	}{
		{
			case1:        []int64{1, 2, 3},
			case2:        []int64{1},
			containsWant: true,
		},
		{
			case1:        []int64{1, 2, 3},
			case2:        []int64{1, 2, 3},
			containsWant: true,
		},
		{
			case1:        []int64{1, 2, 3},
			case2:        []int64{4},
			containsWant: false,
		},
		{
			case1:        []int64{1, 2, 3},
			case2:        []int64{1, 2, 3, 4},
			containsWant: false,
		},
	}
	string64Cases = []struct {
		case1       []string
		case2       []string
		containWant bool
	}{
		{
			case1:       []string{"a", "ab", "abc"},
			case2:       []string{"a"},
			containWant: true,
		},
		{
			case1:       []string{"a", "ab", "abc"},
			case2:       []string{"a", "ab", "abc"},
			containWant: true,
		},
		{
			case1:       []string{"a", "ab", "abc"},
			case2:       []string{"abcd"},
			containWant: false,
		},
		{
			case1:       []string{"a", "ab", "abc"},
			case2:       []string{"a", "ab", "abc", "abcd"},
			containWant: false,
		},
	}
)

func TestContains(t *testing.T) {
	for _, c := range intCases {
		if got := Contains(c.case1, c.case2...); got != c.containWant {
			t.Errorf("containsInt(%v, %v) = %v, want %v", c.case1, c.case2, got, c.containWant)
		}
	}
	for _, c := range int64Cases {
		if got := Contains(c.case1, c.case2...); got != c.containsWant {
			t.Errorf("containsInt(%v, %v) = %v, want %v", c.case1, c.case2, got, c.containsWant)
		}
	}
	for _, c := range string64Cases {
		if got := Contains(c.case1, c.case2...); got != c.containWant {
			t.Errorf("containsInt(%v, %v) = %v, want %v", c.case1, c.case2, got, c.containWant)
		}
	}
}
