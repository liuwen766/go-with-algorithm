package main

import "fmt"

func main() {
	nodeList := getListNode()
	fmt.Print("翻转链表前打印：")
	printListNode(nodeList)
	fmt.Println()

	fmt.Print("翻转链后前打印：")
	printListNode(reverseBetween(nodeList, 2, 4))

}

/**
 * @desc: 给你单链表的头指针 head 和两个整数left 和 right ，其中left <= right 。
 * 请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
 * @data: 2022.1.25 17:30
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var pre, cur, nxt, tmp1, tmp2 *ListNode
	//先确定链表的起始点
	pre = nil
	cur = head
	nxt = head
	i := 0
	for cur != nil && i < left {
		pre = cur
		cur = cur.Next
		i++
	}
	//判断left是否大于链表的长度
	if cur == nil {
		return head
	}
	//开始翻转
	if pre != nil {
		tmp1 = pre // 标记住前一节点和当前节点
		tmp2 = cur
	} else {
		return pre
	}
	for cur != nil && i < right {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
		i++
	}
	//翻转完毕,将链表连起来
	tmp1.Next = pre
	tmp2.Next = nxt
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
