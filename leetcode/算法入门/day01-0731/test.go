package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	res := search(nums, target)
	fmt.Println(res)
}

func search(nums []int, target int) int {
	var arr []int
	fmt.Println(nums[0 : len(nums)/2])
	copy(arr, nums[0:len(nums)/2])
	fmt.Println(arr)
	return 0
}
