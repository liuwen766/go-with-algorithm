package main

import "fmt"

func main() {
	nodeList := getListNode()
	fmt.Print("翻转链表前打印：")
	printListNode(nodeList)
	fmt.Println()

	reverseHead := reverseList(nodeList)
	fmt.Println("迭代翻转链表后的头结点：", reverseHead)

	reverseHead1 := reverse(nodeList)
	fmt.Println("递归翻转链表后的头结点：", reverseHead1)

	fmt.Print("翻转链表后打印：")
	printListNode(reverseHead)
}

func reverse(list *ListNode) *ListNode {
	return nil
}

/**
 * @desc: 迭代翻转链表
 * @data: 2022.1.25 10:05
 */
func reverseList(head *ListNode) *ListNode {
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
	//翻转链表
	return pre
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func printListNode(list *ListNode) {
	for list != nil {
		fmt.Print(list.Val, "—>")
		list = list.Next
	}
}

func getListNode() *ListNode {
	val0 := 1
	val1 := 2
	val2 := 3
	val3 := 4
	val4 := 5
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
