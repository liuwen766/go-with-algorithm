package main

import "fmt"

/**
 * 请实现无重复数字的升序数组的二分查找
 * 给定一个 元素升序的、无重复数字的整型数组 nums 和一个目标值 target ，
 * 写一个函数搜索 nums 中的 target，如果目标值存在返回下标（下标从 0 开始），否则返回 -1
 *
 * @param nums int整型一维数组
 * @param target int整型
 * @return int整型
 */
func search(nums []int, target int) int {
	// write code here
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 5
	res := search(nums, target)
	fmt.Println("二分查找结果:", res)
}
