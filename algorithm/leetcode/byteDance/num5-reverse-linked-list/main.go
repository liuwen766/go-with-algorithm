package main

import "fmt"

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	print(node1)

	newHead := reverseList(node1)

	print(newHead)

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * } 5—>4—>3—>2—>1—>nil
 */
func reverseList(head *ListNode) *ListNode {
	//==========================
	//var pre, cur, nxt *ListNode
	//pre = nil
	//cur = head
	//nxt = head
	//for cur != nil {
	//	nxt = cur.Next
	//	cur.Next = pre
	//	pre = cur
	//	cur = nxt
	//}
	//return pre

	//==========================
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func print(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val, "→")
		node = node.Next
	}
	fmt.Println()
}
