package main

import "fmt"

/**
 * @desc: 反转链表前 N 个节点
 * @data: 2022.1.25 17:37
 */
func main() {
	node := getListNode()
	fmt.Print("翻转链表前打印：")
	printListNode(node)
	fmt.Println()
	n := reverseN(node, 3)
	fmt.Print("翻转链后前打印：")
	printListNode(n)
}

func reverseN(head *ListNode, n int) *ListNode {
	var pre, cur, nxt *ListNode
	pre = nil
	cur = head
	nxt = head
	i := 0
	for cur != nil && i < n {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
		i++
	}
	//翻转完毕,将链表连起来
	head.Next = nxt
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
