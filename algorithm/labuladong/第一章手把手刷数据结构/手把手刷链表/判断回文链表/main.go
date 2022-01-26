package main

import (
	"fmt"
	"github.com/go-with-algorithm/algorithm/common"
)

/**
 * @desc: 判断回文链表
 * @data: 2022.1.25 15:30
 */
func main() {
	nodeList := common.GetListNode()
	fmt.Println("判断回文链表1", isPalindrome1(nodeList))

	fmt.Println("判断回文链表2", isPalindrome2(nodeList))
}

/**
 * @desc: TODO 【快慢指针+反转链表】
 * @data: 2022.1.25 16:36
 */
func isPalindrome2(head *common.ListNode) bool {
	var slow, fast *common.ListNode
	slow = head
	fast = head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	//说明是节点个数是奇数，slow指针还要再跑一次
	if fast != nil {
		slow = slow.Next
	}
	//翻转slow往后的链表
	right := reverse(slow)

	//对比head和right两个链表
	for right != nil {
		if head.Val != right.Val {
			return false
		}
		right = right.Next
		head = head.Next
	}
	return true
}

/**
 * @desc: 递归实现翻转链表
 * @data: 2022.1.25 17:19
 */
func reverse(slow *common.ListNode) *common.ListNode {
	if slow == nil || slow.Next == nil {
		return slow
	}
	//递归
	newHead := reverse(slow.Next)
	slow.Next.Next = slow
	slow.Next = nil
	return newHead
}

/**
 * @desc: 把原始链表反转存入一条新的链表，然后比较这两条链表是否相同
 * @data: 2022.1.25 16:36
 */
func isPalindrome1(head *common.ListNode) bool {
	res1 := printListNode(head)
	//第一步：把链表反转
	var pre, cur, next *common.ListNode
	pre = nil
	cur = head
	next = head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	res2 := printListNode(pre)
	//第二步：对比反转前后的链表
	for i := 0; i < len(res1); i++ {
		if res1[i] != res2[i] {
			return false
		}
	}

	return true
}

func printListNode(list *common.ListNode) []int {
	res := make([]int, 0)
	for list != nil {
		res = append(res, list.Val)
		list = list.Next
	}
	return res
}
