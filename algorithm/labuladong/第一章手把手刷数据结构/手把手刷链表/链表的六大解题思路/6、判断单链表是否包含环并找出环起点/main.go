package main

import "github.com/go-with-algorithm/algorithm/common"

func main() {
	node1 := common.GetListNode()
	node := partition(node1, 2)
	common.PrintListNode(node)
}

func partition(head *common.ListNode, x int) *common.ListNode {
	return nil
}
