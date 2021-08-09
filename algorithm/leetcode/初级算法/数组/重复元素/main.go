package main

import (
	"fmt"
)

func main() {
	var nums = []int{3, 3}
	duplicate := containsDuplicate(nums)
	fmt.Println(nums)
	fmt.Println(duplicate)
}

func containsDuplicate(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	var mymap = make(map[int]struct{})
	for i := range nums {
		if _, ok := mymap[nums[i]]; ok {
			return true
		}
		mymap[nums[i]] = struct{}{}
	}
	return false
}

/**
* @desc: 给定一个整数数组，判断是否存在重复元素。
         如果存在一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。
* @data: 2021.8.9 20:24
*/
//func containsDuplicate(nums []int) bool {
//	if len(nums) == 0 {
//		return false
//	}
//	sort.Ints(nums)
//	for i := 0; i < len(nums)-1; i++ {
//		if nums[i] == nums[i+1] {
//			return true
//		}
//	}
//	return false
//}
