package main

import (
	"fmt"
	"github.com/go-with-algorithm/algorithm/common"
)

func main() {
	root := common.GetTree()
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
	res := levelOrder(root)
	fmt.Print(res)
}

/**
 * @desc: 队列【FIFO】
 * @data: 2022.1.27 21:22
 */
func levelOrder(root *common.TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var queue = []*common.TreeNode{root}
	//计算层级
	//var level int
	for len(queue) > 0 {
		counter := len(queue)
		for 0 < counter {
			counter--
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			res = append(res, queue[0].Val)
			//把第一个节点的左右子节点装完后，排除
			queue = queue[1:]
		}
		//level++
	}
	return res
}

func postOrder(root *common.TreeNode) {
	if root == nil {
		return
	}
	postOrder(root.Left)
	postOrder(root.Right)
	fmt.Print(root.Val)
}

func inOrder(root *common.TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.Left)
	fmt.Print(root.Val)
	inOrder(root.Right)
}

func preOrder(root *common.TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val)
	preOrder(root.Left)
	preOrder(root.Right)
}
