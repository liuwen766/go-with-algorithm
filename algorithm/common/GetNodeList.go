package common

import "fmt"

/**
 * @desc: 获取一个链表：1—>2—>3—>4—>5—>null
 * @data: 2022.2.7 9:41
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintListNode(list *ListNode) {
	for list != nil {
		fmt.Print(list.Val, "—>")
		list = list.Next
	}
}

func GetListNode() *ListNode {
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
