package main

import "fmt"

func main() {
	nums := []int{-2, -1}
	fmt.Println("最大子数组和", maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	var dp = make([]int, len(nums))

	if len(nums) == 1 {
		return nums[0]
	}
	max := nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = getMax(nums[i], dp[i-1]+nums[i])
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}

func getMax(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}
