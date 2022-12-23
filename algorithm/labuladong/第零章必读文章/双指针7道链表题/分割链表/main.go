package main

import (
	"fmt"
	"github.com/go-with-algorithm/algorithm/common"
)

func main() {
	node := common.GetListNode()
	common.PrintListNode(node)
	fmt.Println()
	listNode := partition(node, 3)
	common.PrintListNode(listNode)
}

//给你一个链表的头节点 head 和一个特定值 x ，
//请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
func partition(head *common.ListNode, x int) *common.ListNode {
	dummy1 := &common.ListNode{Val: -1}
	dummy2 := &common.ListNode{Val: -1}
	p := head
	p1 := dummy1
	p2 := dummy2
	for p != nil {
		if p.Val < x {
			p1.Next = p
			p1 = p1.Next
		} else {
			p2.Next = p
			p2 = p2.Next
		}
		// 断开原链表
		tmp := p.Next
		p.Next = nil
		p = tmp
	}
	p1.Next = dummy2.Next
	return dummy1.Next
}
