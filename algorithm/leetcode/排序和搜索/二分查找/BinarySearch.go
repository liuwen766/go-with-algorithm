package main

import (
	"fmt"
)

func main() {

	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 5
	res := search(nums, target)
	index := search1(nums, target)
	fmt.Println(res, index)

}

/**
 * @desc: 二分查找（nums是有序数组）返回下标
 * @data: 2021.8.8 16:38
 */
func search1(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

/**
* @desc: 二分查找（nums是有序数组）
         递归方法只能返回是目标
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
		result = search(arr, target)
	} else {
		arr = nums[len(nums)/2+1:]
		result = search(arr, target)
	}
	return result
}
