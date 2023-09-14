package setx

type Set[T comparable] interface {
	Add(...T)
	Remove(values ...T)
	Contains(values ...T) bool
	Diff(set Set[T]) Set[T]
	Union(set Set[T]) Set[T]
	Intersect(set Set[T]) Set[T]
	IsSubset(set Set[T]) bool
	IsSuperset(set Set[T]) bool
	Slice() []T
	String() string
	Length() int
	Clear()
	Equal(set Set[T]) bool
}
