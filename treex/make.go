package treex

import "errors"

func MakeTree[T any, K comparable](nodes []T, getId func(T) K, getPId func(T) K, rootId K) (*Node[T], error) {
	mapById := make(map[K]T)
	mapByPId := make(map[K][]T)
	for _, node := range nodes {
		id := getId(node)
		pId := getPId(node)
		mapById[id] = node
		mapByPId[pId] = append(mapByPId[pId], node)
	}
	var makeChildren func(id K) []*Node[T]
	makeChildren = func(id K) []*Node[T] {
		elems := mapByPId[id]
		if len(elems) == 0 {
			return nil
		}
		nodes := make([]*Node[T], 0, len(elems))
		for i := range elems {
			nodes = append(nodes, &Node[T]{
				Elem: elems[i],
			})
		}
		for i := range nodes {
			nodes[i].Children = makeChildren(getId(nodes[i].Elem))
		}
		return nodes
	}
	root, ok := mapById[rootId]
	if !ok {
		return nil, errors.New("root id not exist")
	}
	rootNode := Node[T]{
		Elem:     root,
		Children: makeChildren(rootId),
	}
	return &rootNode, nil
}
