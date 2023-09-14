package setx

import (
	"fmt"
)

type hashSet[T comparable] struct {
	hashMap map[T]struct{}
}

func NewHashSet[T comparable](values ...T) Set[T] {
	hashMap := make(map[T]struct{})
	for _, value := range values {
		hashMap[value] = struct{}{}
	}
	return &hashSet[T]{
		hashMap: hashMap,
	}
}

func (s *hashSet[T]) Add(values ...T) {
	for _, value := range values {
		s.hashMap[value] = struct{}{}
	}
}

func (s *hashSet[T]) Remove(values ...T) {
	for _, value := range values {
		delete(s.hashMap, value)
	}
}

func (s *hashSet[T]) Contains(values ...T) bool {
	for _, value := range values {
		if _, ok := s.hashMap[value]; !ok {
			return false
		}
	}
	return true
}

func (s *hashSet[T]) Diff(set Set[T]) Set[T] {
	for key := range set.(*hashSet[T]).hashMap {
		delete(s.hashMap, key)
	}
	return s
}

func (s *hashSet[T]) Union(set Set[T]) Set[T] {
	union := NewHashSet[T]()
	for key := range s.hashMap {
		union.Add(key)
	}
	for key := range set.(*hashSet[T]).hashMap {
		union.Add(key)
	}
	return union
}

func (s *hashSet[T]) Intersect(set Set[T]) Set[T] {
	intersect := NewHashSet[T]()
	for key := range s.hashMap {
		if set.Contains(key) {
			intersect.Add(key)
		}
	}
	return intersect
}

func (s *hashSet[T]) IsSubset(set Set[T]) bool {
	for key := range s.hashMap {
		if !set.Contains(key) {
			return false
		}
	}
	return true
}

func (s *hashSet[T]) IsSuperset(set Set[T]) bool {
	return set.IsSubset(s)
}

func (s *hashSet[T]) Slice() []T {
	slice := make([]T, 0, len(s.hashMap))
	for key := range s.hashMap {
		slice = append(slice, key)
	}
	return slice
}

func (s *hashSet[T]) String() string {
	return fmt.Sprintf("%v", s.Slice())
}

func (s *hashSet[T]) Length() int {
	return len(s.hashMap)
}

func (s *hashSet[T]) Clear() {
	s.hashMap = make(map[T]struct{})
}

func (s *hashSet[T]) Equal(set Set[T]) bool {
	if s.Length() != set.Length() {
		return false
	}
	return s.IsSubset(set)
}
