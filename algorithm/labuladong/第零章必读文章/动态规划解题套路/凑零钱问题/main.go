package main

import (
	"fmt"
	"math"
)

func main() {
	coins := []int{1, 2, 5}
	amount := 6
	fmt.Println("最少的硬币个数【动态规划[该方法会超时]】:", coinChange1(coins, amount))
	fmt.Println("最少的硬币个数【动态数组】:", coinChange2(coins, amount))
}

func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// 数组大小为 amount + 1，初始值也为 amount + 1
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	// base case
	dp[0] = 0
	// 外层 for 循环遍历所有状态的所有取值
	for i := 0; i < len(dp); i++ {
		// 内层 for 循环在求所有选择的最小值
		for _, coin := range coins {
			// 子问题无解，跳过
			if i-coin < 0 {
				continue
			}
			dp[i] = getMin(dp[i], 1+dp[i-coin])
		}
	}
	//初始值没有变化，说明无解
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

/**
 * @desc: 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
 * @data: 2022.1.26 15:00
 */
func coinChange1(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	res := math.MaxInt8
	for i := 0; i < len(coins); i++ {
		subProblem := coinChange1(coins, amount-coins[i])
		//子问题无解则跳过
		if subProblem == -1 {
			continue
		}
		res = getMin(res, 1+subProblem)
	}
	if res == math.MaxInt8 {
		res = -1
	}
	return res
}

func getMin(res int, change1 int) int {
	if res < change1 {
		return res
	}
	return change1
}
