package main

import "github.com/go-with-algorithm/algorithm/common"

func main() {

}

//给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。
//如果两个链表不存在相交节点，返回 null
func getIntersectionNode(headA, headB *common.ListNode) *common.ListNode {
	pA, pB := headA, headB
	for pA != pB {
		if pA != nil {
			pA = pA.Next
			//	链表A走到尾了，开始走链表B
		} else {
			pA = headB
		}
		if pB != nil {
			pB = pB.Next
			//	链表B走到尾了，开始走链表A
		} else {
			pB = headA
		}
	}
	// 走完两个链表，要么找到了相交点pA或者pB都可以返回，要么没找到，返回nil【pA或者pB都是nil】
	return pA
}

// 方法1：双指针一起走完两个链表，如果有相交，则一定相遇
// 方法2：头尾相连，有环说明相交
// 方法3：想办法让双指针到尾部能走一样的长度，这样如果相交一定可以在中间相遇
