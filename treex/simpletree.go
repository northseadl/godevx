package treex

import (
	"errors"
	"fmt"
)

type SimpleTree[T any] struct {
	root  *Node[T]
	getId func(T) interface{}
}

func NewSimpleTree[T any](rootValue T, getId func(T) interface{}) *SimpleTree[T] {
	return &SimpleTree[T]{
		root:  &Node[T]{Elem: rootValue},
		getId: getId,
	}
}

func (t *SimpleTree[T]) Root() *Node[T] {
	return t.root
}

func (t *SimpleTree[T]) Insert(parent, child T) error {
	parentNode := t.Find(parent)
	if parentNode == nil {
		return errors.New("parent node not found")
	}
	parentNode.Children = append(parentNode.Children, &Node[T]{Elem: child})
	return nil
}

func (t *SimpleTree[T]) Remove(node T) error {
	return errors.New("not implemented")
}

func (t *SimpleTree[T]) Find(value T) *Node[T] {
	var found *Node[T]
	var find func(n *Node[T]) bool
	find = func(n *Node[T]) bool {
		if t.getId(n.Elem) == t.getId(value) {
			found = n
			return true
		}
		for _, child := range n.Children {
			if find(child) {
				return true
			}
		}
		return false
	}
	find(t.root)
	return found
}

func (t *SimpleTree[T]) TraversePreOrder() []T {
	var result []T
	var traverse func(n *Node[T])
	traverse = func(n *Node[T]) {
		result = append(result, n.Elem)
		for _, child := range n.Children {
			traverse(child)
		}
	}
	traverse(t.root)
	return result
}

func (t *SimpleTree[T]) TraverseInOrder() []T {
	return nil // Not applicable for a generic tree, only for binary trees
}

func (t *SimpleTree[T]) TraversePostOrder() []T {
	var result []T
	var traverse func(n *Node[T])
	traverse = func(n *Node[T]) {
		for _, child := range n.Children {
			traverse(child)
		}
		result = append(result, n.Elem)
	}
	traverse(t.root)
	return result
}

func (t *SimpleTree[T]) TraverseLevelOrder() []T {
	var result []T
	queue := []*Node[T]{t.root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node.Elem)
		for _, child := range node.Children {
			queue = append(queue, child)
		}
	}
	return result
}

func (t *SimpleTree[T]) Height() int {
	var height func(n *Node[T]) int
	height = func(n *Node[T]) int {
		if n == nil {
			return 0
		}
		maxHeight := 0
		for _, child := range n.Children {
			h := height(child)
			if h > maxHeight {
				maxHeight = h
			}
		}
		return maxHeight + 1
	}
	return height(t.root)
}

func (t *SimpleTree[T]) String() string {
	return fmt.Sprintf("%v", t.TraverseLevelOrder())
}
