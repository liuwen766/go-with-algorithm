package main

import (
	"fmt"
	"github.com/go-with-algorithm/algorithm/common"
)

/**
 * @desc: 反转链表前 N 个节点
 * @data: 2022.1.25 17:37
 */
func main() {
	node := common.GetListNode()
	fmt.Print("翻转链表前打印：")
	common.PrintListNode(node)
	fmt.Println()
	n := reverseN(node, 3)
	fmt.Print("翻转链后前打印：")
	common.PrintListNode(n)
}

func reverseN(head *common.ListNode, n int) *common.ListNode {
	var pre, cur, nxt *common.ListNode
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
