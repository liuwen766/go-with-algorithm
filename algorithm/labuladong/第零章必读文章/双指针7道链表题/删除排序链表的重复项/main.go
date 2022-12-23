package main

import "github.com/go-with-algorithm/algorithm/common"

func main() {
	node := common.GetListNode()
	deleteDuplicates(node)
	common.PrintListNode(node)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *common.ListNode) *common.ListNode {
	// 边界判断
	if head == nil {
		return head
	}
	slow, fast := head, head
	for fast != nil {
		if slow.Val != fast.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	//将slow的后面节点置为null
	slow.Next = nil
	return head
}
