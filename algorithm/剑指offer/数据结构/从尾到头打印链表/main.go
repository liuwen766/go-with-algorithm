package main

import "fmt"

func main() {
	node := &ListNode{
		100,
		&ListNode{
			20,
			nil,
		},
	}
	fmt.Println(printListFromTailToHead(node))
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param head ListNode类
 * @return int整型一维数组
 */
func printListFromTailToHead(head *ListNode) []int {
	// write code here

	var count int
	head1 := head
	for head1 != nil {
		count++
		head1 = head1.Next
	}
	var res []int
	for i := 0; i < count; i++ {
		if head != nil {
			res = append(res, head.Val)
			head = head.Next
		}
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[count-i-1] = res[count-i-1], res[i]
	}
	//for i := count - 1; i >= 0; i-- {
	//	if head != nil {
	//		fmt.Println(head.Val)
	//		res[i] = head.Val
	//		head = head.Next
	//	}
	//	fmt.Println(res)
	//}
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}
