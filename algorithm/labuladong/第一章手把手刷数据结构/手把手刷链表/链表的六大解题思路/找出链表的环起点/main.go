package main

import (
	"fmt"
	"github.com/go-with-algorithm/algorithm/common"
)

func main() {
	nodeList := common.GetListNode()
	fmt.Print("链表里是否有圈：", hasCycle(nodeList))
}

func hasCycle(head *common.ListNode) bool {
	return true
}
