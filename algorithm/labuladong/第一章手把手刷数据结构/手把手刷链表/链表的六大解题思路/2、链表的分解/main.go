package main

import "github.com/go-with-algorithm/algorithm/common"

func main() {
	node1 := common.GetListNode()
	node := partition(node1, 2)
	common.PrintListNode(node)
}

/*
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
*/
func partition(head *common.ListNode, x int) *common.ListNode {
	// 定义两个虚拟头节点
	dummy1 := &common.ListNode{Val: -1}
	dummy2 := &common.ListNode{Val: -1}
	// 定义左右两个指针 left存放小于x的链表，right存放大于x的链表
	left := dummy1
	right := dummy2
	// 定义一个节点用来遍历
	p := head
	for p != nil {
		if p.Val >= x {
			right.Next = p
			right = right.Next
		} else {
			left.Next = p
			left = left.Next
		}
		//断开原链表中的每个节点的 next 指针   todo
		tmp := p.Next
		p.Next = nil
		p = tmp
	}
	//连接两个链表
	left.Next = dummy2.Next

	return dummy1.Next
}
