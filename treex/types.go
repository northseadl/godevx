package treex

type Node[T any] struct {
	Elem     T
	Children []*Node[T]
}
