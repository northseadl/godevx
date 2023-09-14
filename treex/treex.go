package treex

type Tree[T any] interface {
	Root() *Node[T]
	Insert(parent, child T) error
	Remove(node T) error
	Find(value T) *Node[T]
	TraversePreOrder() []T
	TraverseInOrder() []T
	TraversePostOrder() []T
	TraverseLevelOrder() []T
	Height() int
}
