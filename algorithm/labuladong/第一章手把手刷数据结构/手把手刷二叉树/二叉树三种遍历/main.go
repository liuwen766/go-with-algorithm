package main

import (
	"fmt"
	"github.com/go-with-algorithm/algorithm/common"
)

func main() {
	root := getTree()
	//-------------------------------递归遍历------------------------------//
	//前序遍历
	fmt.Print("前序遍历：")
	preOrder(root)
	//中序遍历
	fmt.Println()
	fmt.Print("中序遍历：")
	inOrder(root)
	//后序遍历
	fmt.Println()
	fmt.Print("后序遍历：")
	postOrder(root)

	//-------------------------------迭代遍历------------------------------//
	//前序遍历
	fmt.Println()
	fmt.Print("前序遍历：")
	preOrder(root)
	//中序遍历
	fmt.Println()
	fmt.Print("中序遍历：")
	inOrder(root)
	//后序遍历
	fmt.Println()
	fmt.Print("后序遍历：")
	postOrder(root)
	//-------------------------------第四种遍历：层序遍历------------------------------//
	//层序遍历
	fmt.Println()
	fmt.Print("层序遍历：")
	postOrder(root)
}

func postOrder(root *common.TreeNode) {
	if root != nil {
		postOrder(root.Left)
		postOrder(root.Right)
		fmt.Print(root.Val)
	}
}

func inOrder(root *common.TreeNode) {
	if root != nil {
		inOrder(root.Left)
		fmt.Print(root.Val)
		inOrder(root.Right)
	}
}

func preOrder(root *common.TreeNode) {
	if root != nil {
		fmt.Print(root.Val)
		preOrder(root.Left)
		preOrder(root.Right)
	}

}

/**中序遍历：1234567
       4
    2      6
  1   3  5   7
*/
func getTree() *common.TreeNode {
	var root = new(common.TreeNode)
	var left1 = new(common.TreeNode)
	var right1 = new(common.TreeNode)
	var left2 = new(common.TreeNode)
	var right2 = new(common.TreeNode)
	var left3 = new(common.TreeNode)
	var right3 = new(common.TreeNode)

	root.Val = 4
	left1.Val = 2
	right1.Val = 6
	left2.Val = 1
	right2.Val = 3
	left3.Val = 5
	right3.Val = 7

	root.Left = left1
	root.Right = right1

	left1.Left = left2
	left1.Right = right2

	right1.Left = left3
	right1.Right = right3
	return root
}
