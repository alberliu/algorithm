package algorithm

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	root := InitBTree()
	fmt.Println("先序打印二叉树：")
	PrintBTree(root)
	fmt.Println("二叉树的深度：")
	fmt.Println(GetBTreeDepth(root))
	fmt.Println("两点之间最大距离：")
	GetBTreeMaxDistance(root)
	fmt.Println(maxDistance)
}

type BTreeNode struct {
	Data          int        // 数据
	Left          *BTreeNode // 左子树
	LeftMaxDepth  int        // 左子树最大深度
	Right         *BTreeNode // 右子树
	RightMaxDepth int        // 右子树最大深度
}

// InitBTree 初始化二叉树
func InitBTree() *BTreeNode {
	root := BTreeNode{
		Data: 1,
		Left: &BTreeNode{
			Data: 2,
			Left: &BTreeNode{
				Data:  4,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &BTreeNode{
			Data: 3,
			Left: &BTreeNode{
				Data:  5,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	return &root
}

// PrintBTree 前序遍历二叉树
func PrintBTree(root *BTreeNode) {
	if root != nil {
		fmt.Println(root.Data)
		PrintBTree(root.Left)
		PrintBTree(root.Right)
	}
}

// GetBTreeDepth 获取二叉树的深度
func GetBTreeDepth(node *BTreeNode) int {
	if node == nil {
		return 0
	}
	leftDepth := GetBTreeDepth(node.Left)
	rightDepth := GetBTreeDepth(node.Right)

	if leftDepth >= rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

// GetBTreeMaxDistance 获取两个节点的最大距离
var maxDistance = 0

func GetBTreeMaxDistance(node *BTreeNode) int {
	if node == nil {
		return 0
	}
	leftDepth := GetBTreeDepth(node.Left)
	rightDepth := GetBTreeDepth(node.Right)

	if leftDepth+rightDepth > maxDistance {
		maxDistance = leftDepth + rightDepth
	}

	if leftDepth >= rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}
