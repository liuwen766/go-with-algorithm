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

	fmt.Print("翻转链后前打印：")
	common.PrintListNode(reverseBetween(nodeList, 2, 4))

}

/**
 * @desc: 递归解法
 * @data: 2022/12/11 19:27
 */
func reverseBetween(head *common.ListNode, left int, right int) *common.ListNode {
	if left == 1 {
		return reverseN(head, right)
	}
	// todo
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

var tmpNode *common.ListNode

func reverseN(head *common.ListNode, n int) *common.ListNode {
	if n == 1 {
		tmpNode = head.Next
		return head
	}
	newHead := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = tmpNode
	return newHead
}

/**
 * @desc: 给你单链表的头指针 head 和两个整数left 和 right ，其中left <= right 。
 * 请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
 * @data: 2022.1.25 17:30
 */
func reverseBetween1(head *common.ListNode, left int, right int) *common.ListNode {
	var pre, cur, nxt, tmp1, tmp2 *common.ListNode
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
