package setx

import (
	"testing"
)

func TestHashSet(t *testing.T) {
	set := NewHashSet[int](1, 2, 3)

	if set.Length() != 3 {
		t.Errorf("Expected length 3, got %d", set.Length())
	}

	set.Add(4, 5)

	if !set.Contains(4, 5) {
		t.Error("Expected set to contain 4 and 5")
	}

	set.Remove(1, 2)

	if set.Contains(1, 2) {
		t.Error("Expected set not to contain 1 and 2")
	}

	set2 := NewHashSet[int](3, 4, 5, 6)

	union := set.Union(set2)
	expectedUnion := NewHashSet[int](3, 4, 5, 6)
	if !union.Equal(expectedUnion) {
		t.Errorf("Expected union to be %v, got %v", expectedUnion, union)
	}

	intersect := set.Intersect(set2)
	expectedIntersect := NewHashSet[int](3, 4, 5)
	if !intersect.Equal(expectedIntersect) {
		t.Errorf("Expected intersect to be %v, got %v", expectedIntersect, intersect)
	}

	if !set.IsSubset(union) {
		t.Error("Expected set to be a subset of union")
	}

	if !union.IsSuperset(set) {
		t.Error("Expected union to be a superset of set")
	}

	set.Clear()

	if set.Length() != 0 {
		t.Errorf("Expected length 0 after clear, got %d", set.Length())
	}
}
