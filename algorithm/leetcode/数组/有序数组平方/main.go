package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-7, -3, 2, 3, 11, 12}
	res := sortedSquares(nums)
	fmt.Println(res)
}

/**
 * @desc: 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
 * @data: 2021.8.8 18:32
 */
func sortedSquares(nums []int) []int {
	var ans = make([]int, len(nums))
	//for i := range nums {
	//	nums[i] = nums[i] * nums[i]
	//}
	left := 0
	right := len(nums) - 1
	for pos := len(nums) - 1; pos >= 0; pos-- {
		//if nums[left] > nums[right] {
		//	ans[pos] = nums[left]
		//	left++
		//} else {
		//	ans[pos] = nums[right]
		//	right--
		//}
		if l, r := nums[left]*nums[left], nums[right]*nums[right]; l > r {
			ans[pos] = l
			left++
		} else {
			ans[pos] = r
			right--
		}
	}
	return ans
}

func sortedSquares1(nums []int) []int {
	var ans = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = nums[i] * nums[i]
	}
	sort.Ints(ans)
	return ans
}
