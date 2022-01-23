package main

/**
 * @desc: 生成赫夫曼树
 * step1：排序，从小到大
 * step2：取出两个最小的值作为左右子树节点
 * step3：组成一颗二叉树，根节点为左右子树节点之和
 * step4：将上述结果返回原数组，重复执行step1-step4
 * @data: 2022.1.20 19:40
 */
func main() {

}

type Node struct {
	val   int
	left  *Node
	right *Node
}
