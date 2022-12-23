package main

import "github.com/go-with-algorithm/algorithm/common"

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res int

func maxDepth(root *common.TreeNode) int {
	//边界条件
	if root == nil {
		return 0
	}
	// 获取左子树的最大深度
	resLeft := maxDepth(root.Left)
	// 获取右子树的最大深度
	resRight := maxDepth(root.Right)
	// 获取最终的最大深度【后序遍历】
	res = 1 + getMaxWithTwo(resLeft, resRight)
	//res = getMaxWithThree(res, 1+resLeft, 1+resRight)
	return res
}

func getMaxWithTwo(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func getMaxWithThree(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		} else {
			return c
		}
	} else {
		if b > c {
			return b
		} else {
			return c
		}
	}
}
