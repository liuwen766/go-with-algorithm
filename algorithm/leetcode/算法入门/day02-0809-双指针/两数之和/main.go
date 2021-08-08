package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	res := twoSum(nums, 18)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	hashTable := make(map[int]int)
	for i, val := range nums {
		if p, ok := hashTable[target-val]; ok {
			return []int{p, i}
		}
		hashTable[val] = i
	}
	return []int{}
}

//func twoSum(nums []int, target int) []int {
//	for i := 0; i < len(nums); i++ {
//		for j := i + 1; j < len(nums); j++ {
//			if nums[i]+nums[j] == target {
//				return []int{i, j}
//			}
//		}
//	}
//	return []int{}
//}
