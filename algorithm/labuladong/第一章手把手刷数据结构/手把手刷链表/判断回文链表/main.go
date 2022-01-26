package main

import (
	"fmt"
)

/**
 * @desc: 判断回文链表
 * @data: 2022.1.25 15:30
 */
func main() {
	nodeList := getListNode()
	fmt.Println("判断回文链表1", isPalindrome1(nodeList))

	fmt.Println("判断回文链表2", isPalindrome2(nodeList))
}

/**
 * @desc: TODO 【快慢指针+反转链表】
 * @data: 2022.1.25 16:36
 */
func isPalindrome2(head *ListNode) bool {
	var slow, fast *ListNode
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
func reverse(slow *ListNode) *ListNode {
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
func isPalindrome1(head *ListNode) bool {
	res1 := printListNode(head)
	//第一步：把链表反转
	var pre, cur, next *ListNode
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

func printListNode(list *ListNode) []int {
	res := make([]int, 0)
	for list != nil {
		res = append(res, list.Val)
		list = list.Next
	}
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getListNode() *ListNode {
	val0 := 1
	val1 := 2
	val2 := 3
	val3 := 2
	val4 := 1
	head := new(ListNode)
	node1 := new(ListNode)
	node2 := new(ListNode)
	node3 := new(ListNode)
	node4 := new(ListNode)
	head.Val = val0
	node1.Val = val1
	node2.Val = val2
	node3.Val = val3
	node4.Val = val4
	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	return head
}
