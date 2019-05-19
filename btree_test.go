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
	fmt.Println("按层遍历二叉树")
	PrintBTreeByLayer(root)
	fmt.Println("打印所有路径")
	PrintAllPath(root)
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
			Data: 3,
			Left: &Node{
				Data:  5,
				Left:  nil,
				Right: nil,
			},
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
func GetBTreeDepth(root *Node) int {
	if root == nil {
		return 0
	}
	leftDepth := GetBTreeDepth(root.Left)
	rightDepth := GetBTreeDepth(root.Right)

	if leftDepth >= rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

// GetBTreeMaxDistance 获取两个节点的最大距离
var maxDistance = 0

func GetBTreeMaxDistance(root *Node) int {
	if root == nil {
		return 0
	}
	leftDepth := GetBTreeDepth(root.Left)
	rightDepth := GetBTreeDepth(root.Right)

	if leftDepth+rightDepth > maxDistance {
		maxDistance = leftDepth + rightDepth
	}

	if leftDepth >= rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

// PrintBTreeByLayer 按层打印二叉树
func PrintBTreeByLayer(root *Node) {
	if root == nil {
		return
	}
	nodes := make([]*Node, 0, 10)
	nodes = append(nodes, root)
	index := 0
	for index < len(nodes) {
		fmt.Println(nodes[index].Data)
		if nodes[index].Left != nil {
			nodes = append(nodes, nodes[index].Left)
		}
		if nodes[index].Right != nil {
			nodes = append(nodes, nodes[index].Right)
		}
		index++
	}

}

var (
	nodes = make([]*Node, 10)
	end   = 0
)

// PrintAllPath 打印二叉树根结点到所有叶子节点的路径
func PrintAllPath(root *Node) {
	if root != nil {
		nodes[end] = root
		end++
		if root.Left == nil && root.Right == nil {
			for i := 0; i < end; i++ {
				fmt.Print(nodes[i].Data, "  ")
			}
			fmt.Println()
		}

		PrintAllPath(root.Left)
		PrintAllPath(root.Right)
		end--
	}
}
