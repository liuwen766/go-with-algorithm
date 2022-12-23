package main

import "github.com/go-with-algorithm/algorithm/common"

func main() {

}

//给你一个链表的头节点 head ，判断链表中是否有环。
func hasCycle(head *common.ListNode) bool {
	我, 你 := head, head
	// 你在前面跑，我再后面追啊
	for 你 != nil && 你.Next != nil {
		你 = 你.Next.Next
		我 = 我.Next
		if 你 == 我 {
			return true
		}
	}
	return false
}

//给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
func detectCycle(head *common.ListNode) *common.ListNode {
	我, 你 := head, head
	// 你在前面跑，我再后面追啊
	for 你 != nil && 你.Next != nil {
		你 = 你.Next.Next
		我 = 我.Next
		if 你 == 我 {
			// 停止  相遇说明有环
			break
		}
	}
	// 如果跑完上面的代码之后 还没遇到你
	for 你 == nil || 你.Next == nil {
		return nil
	}
	// 我 再重新开始追一次 这次追上一定在环入口相遇
	我 = head
	for 我 != 你 {
		你 = 你.Next
		我 = 我.Next
	}
	return 我
}
