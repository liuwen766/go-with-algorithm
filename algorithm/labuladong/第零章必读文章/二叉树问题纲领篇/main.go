package main

import (
	"fmt"
	"github.com/go-with-algorithm/algorithm/common"
)

/**
 * @desc: 遇到一道二叉树的题目时的通用思考过程是：
 * 是否可以通过遍历一遍二叉树得到答案？如果不能的话，是否可以定义一个递归函数，通过子问题（子树）的答案推导出原问题的答案？
 * @data: 2022.1.27 20:59
 */
func main() {

	fmt.Println("数组的递归遍历")
	arr := []int{1, 2, 3, 4, 5}
	traverseArr(arr, 0)
	fmt.Println()

	fmt.Println("链表的递归遍历")
	node := common.GetListNode()
	traverseList(node)
	fmt.Println()

	fmt.Println("二叉树的递归遍历")
	root := common.GetTree()
	traverseTree(root)
	fmt.Println()

	//二叉树的所有问题，就是让你在前中后序位置注入巧妙的代码逻辑，去达到自己的目的。
	//二叉树题目的递归解法可以分两类思路，第一类是遍历一遍二叉树得出答案，第二类是通过分解问题计算出答案，
	//这两类思路分别对应着 回溯算法核心框架 和 动态规划核心框架。

	//以获取二叉树的最大深度为例
	//是否可以通过遍历一遍二叉树得到答案？
	depth1 := getMaxDepthWithFun1(root)
	//是否可以定义一个递归函数，通过子问题（子树）的答案推导出原问题的答案？
	depth2 := getMaxDepthWithFun2(root)
	fmt.Println("二叉树的最大深度：", depth1)
	fmt.Println("二叉树的最大深度：", depth2)

	//前序位置的代码执行是自顶向下的，而后序位置的代码执行是自底向上的
	//意味着前序位置的代码只能从函数参数中获取父节点传递来的数据，而后序位置的代码不仅可以获取参数数据，
	//还可以获取到子树通过函数返回值传递回来的数据。
	//考虑两个问题：
	//1、如果把根节点看做第 1 层，如何打印出每一个节点所在的层数？【前序遍历】
	//2、如何打印出每个节点的左右子树各有多少节点？【后序遍历】
	//获取二叉树的最大路径 可以前序遍历，可以后序遍历，但后序遍历更好
	fmt.Println("获取二叉树的最大路径", diameterOfBinaryTree(root))

}

var maxDiameter int

func diameterOfBinaryTree(root *common.TreeNode) int {
	getMaxDepth(root)
	return maxDiameter
}

func getMaxDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := getMaxDepth(root.Left)
	rightMax := getMaxDepth(root.Right)
	myDiameter := leftMax + rightMax
	maxDiameter = getMax(maxDiameter, myDiameter)
	return 1 + getMax(leftMax, rightMax)
}

//第二类是通过分解问题计算出答案
func getMaxDepthWithFun2(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := getMaxDepthWithFun2(root.Left)
	rightMax := getMaxDepthWithFun2(root.Right)
	/// 整棵树的最大深度等于左右子树的最大深度取最大值,然后再加上根节点自己
	maxDepth := getMax(leftMax, rightMax) + 1
	return maxDepth
}

func getMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

var ans int
var depth int

//第一类是遍历一遍二叉树得出答案
func getMaxDepthWithFun1(root *common.TreeNode) int {
	traverse(root)
	return ans
}

func traverse(root *common.TreeNode) {
	if root == nil {
		ans = getMax(ans, depth)
	}
	depth++
	if root != nil {
		traverse(root.Left)
		traverse(root.Right)
	}
	depth--
}

var res []int

/**
 * @desc: 二叉树的递归遍历
 * @data: 2022.1.27 16:47
 */
func traverseTree(root *common.TreeNode) {
	if root == nil {
		return
	}
	//前序遍历
	res = append(res, root.Val)
	fmt.Print("—>", root.Val)
	traverseTree(root.Left)
	//中序遍历
	//fmt.Print("—>", root.Val)
	traverseTree(root.Right)
	////后序遍历
	//fmt.Print("—>", root.Val)
}

/**
 * @desc: 链表的递归遍历
 * @data: 2022.1.27 16:30
 */
func traverseList(node *common.ListNode) {
	if node == nil {
		return
	}
	//前序遍历
	fmt.Print("—>", node.Val)
	traverseList(node.Next)
	//后序遍历
	//fmt.Print("—>", node.Val)
}

/**
 * @desc: 数组的递归遍历
 * @data: 2022.1.27 16:36
 */
func traverseArr(arr []int, index int) {
	if index == len(arr) {
		return
	}
	//前序遍历
	fmt.Print(">", arr[index])
	traverseArr(arr, index+1)
	//后序遍历
	//fmt.Print(">", arr[index])
}
