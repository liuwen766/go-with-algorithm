package main

import (
	"fmt"
	"github.com/go-with-algorithm/algorithm/common"
)

func main() {
	/**中序遍历结果：1234567
	       4
	    2      6
	  1   3  5   7
	*/
	root := common.GetTree()
	// 中序遍历
	inTraverse(root)
}

// 中序遍历
func inTraverse(root *common.TreeNode) {
	if root == nil {
		return
	}
	inTraverse(root.Left)
	fmt.Print(root.Val)
	inTraverse(root.Right)
}
