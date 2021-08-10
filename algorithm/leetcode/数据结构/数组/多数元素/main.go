package main

import (
	"fmt"
)

func main() {
	var nums = []int{3}
	fmt.Println(majorityElement(nums))

}

//func majorityElement(nums []int) int {
//	//可以假设数组是非空的，并且给定的数组总是存在多数元素。
//	sort.Ints(nums)
//	return nums[len(nums)/2]
//}

/**
 * @desc: 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素
 * @data: 2021.8.10 20:25
 */
func majorityElement(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	var mapInt = make(map[int]int)
	for _, val := range nums {
		if _, has := mapInt[val]; has {

			mapInt[val] = mapInt[val] + 1
			if mapInt[val] > n/2 {
				return val
			}

		} else {
			mapInt[val] = 1
		}
	}
	return 0
}
