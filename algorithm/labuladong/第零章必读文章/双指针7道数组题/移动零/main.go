package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12, 0}
	moveZeroes(nums)
	fmt.Println(nums)
}

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// [2, 0, 1, 0, 3, 12,0]
func moveZeroes(nums []int) {
	//移除0元素即可
	slow, fast := 0, 0
	//先移除0元素
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	//再将移除0元素后的数组 slow位置之后的数置0
	for slow < len(nums) {
		nums[slow] = 0
		slow++
	}
}
