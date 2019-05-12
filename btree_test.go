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
}

type Node struct {
	Data          int   // 数据
	Left          *Node // 左子树
	LeftMaxDepth  int   // 左子树最大深度
	Right         *Node // 右子树
	RightMaxDepth int   // 右子树最大深度
}

func InitBTree() *Node {
	root := Node{
		Data: 1,
		Left: &Node{
			Data: 2,
			Left: &Node{
				Data:  4,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &Node{
			Data:  3,
			Left:  nil,
			Right: nil,
		},
	}
	return &root
}

func PrintBTree(root *Node) {
	if root != nil {
		fmt.Println(root.Data)
		PrintBTree(root.Left)
		PrintBTree(root.Right)
	}
}

// GetBTreeDepth
func GetBTreeDepth(node *Node) int {
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
