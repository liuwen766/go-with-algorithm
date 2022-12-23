package main

import "github.com/go-with-algorithm/algorithm/common"

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
