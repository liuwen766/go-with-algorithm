package main

import "fmt"

/**
 * @desc: 二叉树节点定义
 * @data: 2022.1.25 14:08
 */
type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

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

func postOrder(root *TreeNode) {
	if root != nil {
		postOrder(root.left)
		postOrder(root.right)
		fmt.Print(root.val)
	}
}

func inOrder(root *TreeNode) {
	if root != nil {
		inOrder(root.left)
		fmt.Print(root.val)
		inOrder(root.right)
	}
}

func preOrder(root *TreeNode) {
	if root != nil {
		fmt.Print(root.val)
		preOrder(root.left)
		preOrder(root.right)
	}

}

/**中序遍历：1234567
       4
    2      6
  1   3  5   7
*/
func getTree() *TreeNode {
	var root = new(TreeNode)
	var left1 = new(TreeNode)
	var right1 = new(TreeNode)
	var left2 = new(TreeNode)
	var right2 = new(TreeNode)
	var left3 = new(TreeNode)
	var right3 = new(TreeNode)

	root.val = 4
	left1.val = 2
	right1.val = 6
	left2.val = 1
	right2.val = 3
	left3.val = 5
	right3.val = 7

	root.left = left1
	root.right = right1

	left1.left = left2
	left1.right = right2

	right1.left = left3
	right1.right = right3
	return root
}
