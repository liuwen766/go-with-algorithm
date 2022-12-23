package main

import "github.com/go-with-algorithm/algorithm/common"

func main() {
	node1 := common.GetListNode()
	nodes := []*common.ListNode{node1, node1, node1}
	node := mergeKLists(nodes)
	common.PrintListNode(node)

	//node2 := common.GetListNodeNil()
	//node3 := common.GetListNode()
	//noderes := merge2Lists(node2, node3)
	//common.PrintListNode(noderes)
}

/*
给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表
*/
func mergeKLists(lists []*common.ListNode) *common.ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	res := lists[0]
	// 两两合并
	for i := 1; i < len(lists); i++ {
		res = merge2Lists(res, lists[i])
	}
	return res
}

func merge2Lists(node1 *common.ListNode, node2 *common.ListNode) *common.ListNode {
	dummy := &common.ListNode{Val: -1}
	left := node1
	right := node2
	p := dummy
	for left != nil && right != nil {
		if left.Val >= right.Val {
			p.Next = right
			right = right.Next
		} else {
			p.Next = left
			left = left.Next
		}
		p = p.Next
	}
	if left == nil {
		p.Next = right
	}
	if right == nil {
		p.Next = left
	}

	return dummy.Next
}
