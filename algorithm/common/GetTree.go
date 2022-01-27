package common

/**
 * @desc: 二叉树节点定义
 * @data: 2022.1.25 14:08
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**中序遍历：1234567
       4
    2      6
  1   3  5   7
*/
func GetTree() *TreeNode {
	var root = new(TreeNode)
	var left1 = new(TreeNode)
	var right1 = new(TreeNode)
	var left2 = new(TreeNode)
	var right2 = new(TreeNode)
	var left3 = new(TreeNode)
	var right3 = new(TreeNode)

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
