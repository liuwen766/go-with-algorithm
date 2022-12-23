package main

import "github.com/go-with-algorithm/algorithm/common"

func main() {

}

//给定一个头结点为 head 的非空单链表，返回链表的中间结点。
func middleNode(head *common.ListNode) *common.ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
