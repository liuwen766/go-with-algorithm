package main

import (
	"fmt"
	"github.com/go-with-algorithm/algorithm/common"
)

func main() {
	node := common.GetListNode()
	common.PrintListNode(node)
	fmt.Println()
	end := removeNthFromEnd(node, 2)
	common.PrintListNode(end)
}

func removeNthFromEnd(head *common.ListNode, n int) *common.ListNode {
	//使用了虚拟头结点的技巧，也是为了防止出现空指针的情况，比如说链表总共有 5 个节点，题目就让你删除倒数第 5 个节点，
	//也就是第一个节点，那按照算法逻辑，应该首先找到倒数第 6 个节点。但第一个节点前面已经没有节点了，这就会出错。
	dummy := &common.ListNode{Val: -1}
	dummy.Next = head
	// 双指针
	// 快指针先走n步
	fast := head
	for i := 0; i < n && fast != nil; i++ {
		fast = fast.Next
	}
	// 快慢指针一起走
	slow := dummy
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	// 断开连接
	slow.Next = slow.Next.Next
	return dummy.Next
}

func removeNthFromEnd1(head *common.ListNode, n int) *common.ListNode {
	dummy := &common.ListNode{Val: -1}
	// 双指针
	dummy.Next = head
	find := getKthFromEnd(head, n+1)
	find.Next = find.Next.Next
	return dummy.Next
}

func getKthFromEnd(head *common.ListNode, k int) *common.ListNode {
	//双指针
	//先让第一个指针跑k个节点
	fast := head
	slow := head
	for i := 0; i < k && fast != nil; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
