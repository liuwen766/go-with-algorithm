package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	//定义两个指针
	left, right := 0, len(numbers)-1
	for left < right {
		if numbers[left]+numbers[right] > target {
			right--
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			return []int{left + 1, right + 1}
		}
	}
	return []int{left + 1, right + 1}
}

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5, 6, 8, 19, 25}, 14))
}

// https://leetcode.cn/problems/he-wei-sde-liang-ge-shu-zi-lcof/
func twoSum1(nums []int, target int) []int {
	//定义两个指针
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] > target {
			right--
		} else if nums[left]+nums[right] < target {
			left++
		} else {
			return []int{nums[left], nums[right]}
		}
	}
	return []int{nums[left], nums[right]}
}

// https://leetcode.cn/problems/kLl5u1/
func twoSum2(numbers []int, target int) []int {
	//定义两个指针
	left, right := 0, len(numbers)-1
	for left < right {
		if numbers[left]+numbers[right] > target {
			right--
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			return []int{left, right}
		}
	}
	return []int{left, right}
}
