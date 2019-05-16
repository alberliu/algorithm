package algorithm

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	head := InitLinkedList()
	PritfLinkedList(head)
}

type LinkedListNode struct {
	Data int
	Next *LinkedListNode
}

// InitLinkedList 初始化链表
func InitLinkedList() *LinkedListNode {
	head := &LinkedListNode{
		Next: &LinkedListNode{
			Data: 1,
			Next: &LinkedListNode{
				Data: 3,
				Next: &LinkedListNode{
					Data: 4,
					Next: &LinkedListNode{
						Data: 5,
						Next: nil,
					},
				},
			},
		},
	}
	return head
}

func PritfLinkedList(head *LinkedListNode) {
	for head.Next != nil {
		head = head.Next
		fmt.Println(head.Data)
	}
}
