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
	// 后序遍历
	postTraverse(root)
}

// 后序遍历
func postTraverse(root *common.TreeNode) {
	if root == nil {
		return
	}
	postTraverse(root.Left)
	postTraverse(root.Right)
	fmt.Print(root.Val)
}
