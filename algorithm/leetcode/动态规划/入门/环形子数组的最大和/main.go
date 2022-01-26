package main

import "fmt"

func main() {
	nums := []int{-5, -3, -5}
	fmt.Println("最大子数组和", maxSubarraySumCircular(nums))
}

/**
 * @desc: 给定一个由整数数组 A 表示的环形数组 C，求 C 的非空子数组的最大可能和。
 * 子数组最多只能包含固定缓冲区 A 中的每个元素一次
 * 思路：求max【最大子序列和，sum[nums]-最小子序列和】
 * @data: 2022.1.26 10:06
 */
func maxSubarraySumCircular(nums []int) int {
	sum := getSum(nums)
	maxSum := maxSubArray(nums)
	minSum := minSubArray(nums)
	summinuxmin := sum - minSum
	//一种特殊情况是数组全为负数，也就是SUM-最小子序和==0，最大子序和等于数组中最大的那个
	if sum-minSum == 0 {
		summinuxmin = maxSum
	}
	return getMax(maxSum, summinuxmin)
}

func getSum(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

func minSubArray(nums []int) int {
	dp := make([]int, len(nums))
	if len(nums) == 1 {
		return nums[0]
	}
	min := nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = getMin(dp[i], dp[i-1]+nums[i])
		if min > dp[i] {
			min = dp[i]
		}
	}
	return min
}

func getMin(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
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

func getMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
