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
	// 前序遍历
	preTraverse(root)
}

// 前序遍历
func preTraverse(root *common.TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val)
	preTraverse(root.Left)
	preTraverse(root.Right)
}

// 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
// 需要递归，需要返回值，必须定义一个全局变量
var res []int

func preorderTraversal(root *common.TreeNode) []int {
	res = make([]int, 0)
	// 如果自己搞不定，那么定义一个递归函数
	preTraversal(root)
	return res
}

// 前序遍历
func preTraversal(root *common.TreeNode) {
	if root == nil {
		return
	}
	res = append(res, root.Val)
	preTraversal(root.Left)
	preTraversal(root.Right)
}
