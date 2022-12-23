package main

import "github.com/go-with-algorithm/algorithm/common"

// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
func mergeTwoLists(list1 *common.ListNode, list2 *common.ListNode) *common.ListNode {
	// 一个虚拟头结点
	dummy := &common.ListNode{Val: -1}
	// 定义一个节点用来遍历
	p := dummy
	// 定义左右两个指针
	left := list1
	right := list2
	// 比大小，升序排列
	for left != nil && right != nil {
		if left.Val >= right.Val {
			// 右边的值小，右指针移
			p.Next = right
			right = right.Next
		} else {
			// 左边指针移
			p.Next = left
			left = left.Next
		}
		p = p.Next
	}
	// 考虑有一个链表为null的情况
	if left == nil {
		p.Next = right
	}
	if right == nil {
		p.Next = left
	}
	// 直接返回虚拟节点的下一个节点即可
	return dummy.Next
}
