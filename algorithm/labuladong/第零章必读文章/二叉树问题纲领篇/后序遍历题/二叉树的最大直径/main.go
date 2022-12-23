package 二叉树的最大直径

import "github.com/go-with-algorithm/algorithm/common"

//给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。
//这条路径可能穿过也可能不穿过根结点。
// TODO 未完成ac
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var maxDiameter int

func diameterOfBinaryTree(root *common.TreeNode) int {
	// 由最大深度获取最大直径
	getMaxDepth(root)
	return maxDiameter
}

func getMaxDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := getMaxDepth(root.Left)
	rightDepth := getMaxDepth(root.Right)
	// 更新最大直径
	diameter := leftDepth + rightDepth
	maxDiameter = getMaxWithTwo(maxDiameter, diameter)
	return 1 + getMaxWithTwo(leftDepth, rightDepth)
}

func getMaxWithTwo(a, b int) int {
	if a > b {
		return a
	}
	return b
}
