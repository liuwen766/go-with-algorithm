package main

import (
	"fmt"
)

func main() {

	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	res := search(nums, target)
	fmt.Println(res)

}

/**
 * @desc: 二分查找（nums是个有序数组）
 * @data: 2021.7.31 0:00
 */
func search(nums []int, target int) int {
	var result int
	var arr []int
	if nums == nil || len(nums) == 0 {
		return result
	}
	//1.递归结束条件
	if nums[len(nums)/2] == target {
		result = nums[len(nums)/2]
		return result
	}
	//2.分解数组
	if nums[len(nums)/2] > target {
		arr = nums[0 : len(nums)/2]
		search(arr, target)
	} else {
		arr = nums[len(nums)/2+1:]
		search(arr, target)
	}
	return result
}
