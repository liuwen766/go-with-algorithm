package main

import "fmt"

/*
给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，
使每个元素 只出现一次 ，返回删除后数组的新长度。
元素的 相对顺序 应该保持 一致 。
*/
func removeDuplicates(nums []int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		// 递增数组 {1, 2, 2, 3, 4, 5, 5, 6, 7, 8, 9}
		if nums[fast] > nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

func main() {
	nums := []int{1, 2, 2, 3, 4, 5, 5, 6, 7, 8, 9}
	length := removeDuplicates(nums)
	fmt.Println("数组的新长度:", length)
	fmt.Println("新数组:", nums)

	val := 5
	numsNew := []int{1, 2, 2, 3, 4, 5, 5, 6, 7, 8, 9}
	lengthNew := removeElement(numsNew, val)
	fmt.Println("数组的新长度:", lengthNew)
	fmt.Println("新数组:", numsNew)
}

func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow + 1
}
