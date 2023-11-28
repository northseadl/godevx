package treex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNAryTree(t *testing.T) {
	tree := &SimpleTree[int]{}

	// 测试设置根节点
	root := &SimpleTreeNode[int]{value: 1}
	tree.SetRoot(root)
	assert.Equal(t, root, tree.Root())

	// 测试插入节点
	err := tree.Insert(1, 2)
	assert.Nil(t, err)
	err = tree.Insert(1, 3)
	assert.Nil(t, err)
	err = tree.Insert(2, 4)
	assert.Nil(t, err)
	err = tree.Insert(2, 5)
	assert.Nil(t, err)

	// 测试查找节点
	node := tree.Find(3)
	assert.NotNil(t, node)
	assert.Equal(t, 3, node.Value())

	// 测试遍历
	preOrderNodes := tree.Traverse(PreOrder)
	assert.Equal(t, []int{1, 2, 4, 5, 3}, extractValues(preOrderNodes))

	postOrderNodes := tree.Traverse(PostOrder)
	assert.Equal(t, []int{4, 5, 2, 3, 1}, extractValues(postOrderNodes))

	levelOrderNodes := tree.Traverse(LevelOrder)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, extractValues(levelOrderNodes))

	// 测试移除节点
	err = tree.Remove(2)
	assert.Nil(t, err)
	assert.Nil(t, tree.Find(2))
	assert.Nil(t, tree.Find(4))
	assert.Nil(t, tree.Find(5))

	// 测试移除根节点失败
	err = tree.Remove(1)
	assert.NotNil(t, err)
}

func extractValues(nodes []*SimpleTreeNode[int]) []int {
	values := make([]int, len(nodes))
	for i, node := range nodes {
		values[i] = node.Value()
	}
	return values
}

func TestMakeSimpleTree(t *testing.T) {
	data := []string{"A", "B", "C", "D", "E", "F", "G"}
	relationFn := func(item string) (parentValue string, value string) {
		switch item {
		case "A":
			return "A", "A"
		case "B", "C":
			return "A", item
		case "D", "E":
			return "B", item
		case "F", "G":
			return "C", item
		default:
			return "", ""
		}
	}

	tree, err := MakeSimpleTree(data, relationFn)
	assert.NoError(t, err)

	expectedPreOrder := []string{"A", "B", "D", "E", "C", "F", "G"}
	preOrder := tree.Traverse(PreOrder)
	for i, node := range preOrder {
		assert.Equal(t, expectedPreOrder[i], node.Value())
	}
}

func TestMakeSimpleForest(t *testing.T) {
	data := []string{"A", "B", "C", "D", "E", "F", "G"}
	relationFn := func(item string) (parentValue string, value string) {
		switch item {
		case "A", "B", "C":
			return "", item
		case "D", "E":
			return "B", item
		case "F", "G":
			return "C", item
		default:
			return "", ""
		}
	}

	forest, err := MakeSimpleForest(data, relationFn)
	assert.NoError(t, err)

	expectedPreOrders := [][]string{
		{"A"},
		{"B", "D", "E"},
		{"C", "F", "G"},
	}

	for i, tree := range forest {
		preOrder := tree.Traverse(PreOrder)
		for j, node := range preOrder {
			assert.Equal(t, expectedPreOrders[i][j], node.Value())
		}
	}
}
