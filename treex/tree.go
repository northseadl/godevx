package treex

import (
	"fmt"
)

type SimpleTreeNode[T comparable] struct {
	value    T
	parent   *SimpleTreeNode[T]
	children []*SimpleTreeNode[T]
}

func (n *SimpleTreeNode[T]) Value() T {
	return n.value
}

func (n *SimpleTreeNode[T]) SetValue(value T) {
	n.value = value
}

func (n *SimpleTreeNode[T]) Parent() *SimpleTreeNode[T] {
	if n.parent == nil {
		return nil
	}
	return n.parent
}

func (n *SimpleTreeNode[T]) SetParent(node *SimpleTreeNode[T]) {
	n.parent = node
}

func (n *SimpleTreeNode[T]) Children() []*SimpleTreeNode[T] {
	nodes := make([]*SimpleTreeNode[T], len(n.children))
	for i, child := range n.children {
		nodes[i] = child
	}
	return nodes
}

func (n *SimpleTreeNode[T]) AddChild(node *SimpleTreeNode[T]) {
	node.SetParent(n)
	n.children = append(n.children, node)
}

func (n *SimpleTreeNode[T]) RemoveChild(node SimpleTreeNode[T]) {
	for i, child := range n.children {
		if child.value == node.value {
			n.children = append(n.children[:i], n.children[i+1:]...)
			break
		}
	}
}

type SimpleTree[T comparable] struct {
	root *SimpleTreeNode[T]
}

func (t *SimpleTree[T]) Root() *SimpleTreeNode[T] {
	return t.root
}

func (t *SimpleTree[T]) SetRoot(node *SimpleTreeNode[T]) {
	t.root = node
}

func (t *SimpleTree[T]) Find(value T) *SimpleTreeNode[T] {
	return find(t.root, value)
}

func find[T comparable](node *SimpleTreeNode[T], value T) *SimpleTreeNode[T] {
	if node == nil {
		return nil
	}
	if node.value == value {
		return node
	}
	for _, child := range node.children {
		if found := find(child, value); found != nil {
			return found
		}
	}
	return nil
}

func (t *SimpleTree[T]) Insert(parentValue, value T) error {
	parent := t.Find(parentValue)
	if parent == nil {
		return fmt.Errorf("parent not found")
	}
	node := &SimpleTreeNode[T]{value: value}
	parent.AddChild(node)
	return nil
}

func (t *SimpleTree[T]) Remove(value T) error {
	node := t.Find(value)
	if node == nil {
		return fmt.Errorf("node not found")
	}
	parent := node.Parent()
	if parent == nil {
		return fmt.Errorf("cannot remove root")
	}
	parent.RemoveChild(*node)
	return nil
}

func (t *SimpleTree[T]) Traverse(traversalType TraversalType) []*SimpleTreeNode[T] {
	switch traversalType {
	case PreOrder:
		return preOrder(t.root)
	case PostOrder:
		return postOrder(t.root)
	case LevelOrder:
		return levelOrder(t.root)
	default:
		return nil
	}
}

func preOrder[T comparable](node *SimpleTreeNode[T]) []*SimpleTreeNode[T] {
	if node == nil {
		return nil
	}
	nodes := []*SimpleTreeNode[T]{node}
	for _, child := range node.children {
		nodes = append(nodes, preOrder(child)...)
	}
	return nodes
}

func postOrder[T comparable](node *SimpleTreeNode[T]) []*SimpleTreeNode[T] {
	if node == nil {
		return nil
	}
	var nodes []*SimpleTreeNode[T]
	for _, child := range node.children {
		nodes = append(nodes, postOrder(child)...)
	}
	nodes = append(nodes, node)
	return nodes
}

func levelOrder[T comparable](node *SimpleTreeNode[T]) []*SimpleTreeNode[T] {
	if node == nil {
		return nil
	}
	nodes := []*SimpleTreeNode[T]{node}
	for i := 0; i < len(nodes); i++ {
		for _, child := range nodes[i].children {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func MakeSimpleTree[T comparable](list []T, relationFn func(T) (parentValue T, value T)) (*SimpleTree[T], error) {
	tree := &SimpleTree[T]{}
	nodes := make(map[T]*SimpleTreeNode[T])

	for _, item := range list {
		parentValue, value := relationFn(item)
		node, exists := nodes[value]
		if !exists {
			node = &SimpleTreeNode[T]{value: value}
			nodes[value] = node
		}

		if parentValue == value {
			if tree.root != nil {
				return nil, fmt.Errorf("multiple root nodes")
			}
			tree.root = node
		} else {
			parent, exists := nodes[parentValue]
			if !exists {
				parent = &SimpleTreeNode[T]{value: parentValue}
				nodes[parentValue] = parent
			}
			parent.AddChild(node)
		}
	}

	if tree.root == nil {
		return nil, fmt.Errorf("no root node")
	}

	return tree, nil
}

func MakeSimpleForest[T comparable](list []T, relationFn func(T) (parentValue T, value T)) ([]*SimpleTree[T], error) {
	var forest []*SimpleTree[T]
	nodes := make(map[T]*SimpleTreeNode[T])
	for _, item := range list {
		parentValue, value := relationFn(item)
		node, exists := nodes[value]
		if !exists {
			node = &SimpleTreeNode[T]{value: value}
			nodes[value] = node
		}
		if parentValue == *new(T) {
			tree := &SimpleTree[T]{root: node}
			forest = append(forest, tree)
		} else {
			parent, exists := nodes[parentValue]
			if !exists {
				parent = &SimpleTreeNode[T]{value: parentValue}
				nodes[parentValue] = parent
			}
			parent.AddChild(node)
		}
	}
	return forest, nil
}
