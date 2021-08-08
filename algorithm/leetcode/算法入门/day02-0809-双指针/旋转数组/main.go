package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 2)
	//revise(nums)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	revise(nums)
	i := k % len(nums)
	revise(nums[:i])
	revise(nums[i:])
}

//反转数组
func revise(nums []int) {
	for i, j := 0, len(nums)-1; i < j/2; i++ {
		nums[i], nums[j-i] = nums[j-i], nums[i]
	}
}

//反转数组
func revise1(nums []int) {
	var newArr = make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		newArr[len(nums)-1-i] = nums[i]
	}
	copy(nums, newArr)
}

func rotate2(nums []int, k int) {
	var newArr = make([]int, len(nums))
	for i, val := range nums {
		newArr[(i+k)%len(nums)] = val
	}
	//nums=newArr
	copy(nums, newArr)
}

func rotate1(nums []int, k int) {
	newArr := nums
	for i := 0; i < len(nums); i++ {
		newArr = append(newArr, nums[i])
	}
	var k2 = len(nums) - (k % len(nums))
	for i := k2; i < k2+len(nums); i++ {
		nums[i-k2] = newArr[i]
	}
}
