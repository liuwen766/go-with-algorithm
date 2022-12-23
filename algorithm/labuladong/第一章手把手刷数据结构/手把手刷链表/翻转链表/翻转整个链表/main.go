package main

import (
	"fmt"
	"github.com/go-with-algorithm/algorithm/common"
)

func main() {
	nodeList := common.GetListNode()
	fmt.Print("翻转链表前打印：")
	common.PrintListNode(nodeList)
	fmt.Println()

	reverseHead := reverseList(nodeList)
	fmt.Println("迭代翻转链表后的头结点：", reverseHead)

	reverseHead1 := reverseList1(nodeList)
	fmt.Println("递归翻转链表后的头结点：", reverseHead1)

	fmt.Print("翻转链表后打印：")
	common.PrintListNode(reverseHead)
}

/**
* @desc: 递归
* @data: 2022/12/11 18:44
 */
func reverseList1(head *common.ListNode) *common.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

/**
 * @desc: 迭代翻转链表
 * @data: 2022.1.25 10:05
 */
func reverseList(head *common.ListNode) *common.ListNode {
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
	//翻转链表
	return pre
}
