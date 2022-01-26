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
	return -1
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
