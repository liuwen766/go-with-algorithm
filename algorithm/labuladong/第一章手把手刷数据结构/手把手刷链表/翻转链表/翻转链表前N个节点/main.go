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
	fmt.Println()
	newHead := reverseN2(node, 3)
	fmt.Print("翻转链后前打印：")
	common.PrintListNode(newHead)
}

/**
* @desc: 递归解法
* @data: 2022/12/11 19:18
  1→2→3→4→5→nil
  3→2→1→4→5→nil
*/
var tmpNext *common.ListNode

func reverseN2(head *common.ListNode, n int) *common.ListNode {
	if n == 1 {
		//记录下第 n + 1 个节点
		tmpNext = head.Next
		return head
	}
	newHead := reverseN2(head.Next, n-1)
	head.Next.Next = head
	// 让反转之后的 head 节点和后面的节点连起来
	head.Next = tmpNext
	return newHead
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
