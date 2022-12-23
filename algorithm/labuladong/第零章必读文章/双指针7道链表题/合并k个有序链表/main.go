package main

import "github.com/go-with-algorithm/algorithm/common"

func main() {

}

func mergeKLists(lists []*common.ListNode) *common.ListNode {
	if len(lists) == 0 {
		return nil
	}
	res := lists[0]
	for i := 1; i < len(lists); i++ {
		res = merge2Lists(res, lists[i])
	}
	return res
}

func merge2Lists(node1 *common.ListNode, node2 *common.ListNode) *common.ListNode {
	dummy := &common.ListNode{Val: -1}
	p1 := node1
	p2 := node2
	p := dummy
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p.Next = p1
			p = p.Next
			p1 = p1.Next
		} else {
			p.Next = p2
			p = p.Next
			p2 = p2.Next
		}
	}
	if p1 == nil {
		p.Next = p2
	}
	if p2 == nil {
		p.Next = p1
	}
	return dummy.Next
}
