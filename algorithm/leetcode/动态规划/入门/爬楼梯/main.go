package main

import (
	"fmt"
)

func main() {
	fmt.Println("共有多少种爬楼梯方法:", climbStairs(4))

	fmt.Println("使用最小花费爬楼梯共有多少种爬楼梯方法:", minCostClimbingStairs([]int{10, 13, 20}))
}

func climbStairs(n int) int {
	// 超出时间限制
	//if n <= 2 {
	//	return n
	//}
	////  楼梯:0 1 2 3 4 5 6
	////  方法:0 1 2 3 5 8 13
	//return climbStairs(n-1) + climbStairs(n-2)
	if n <= 2 {
		return n
	}
	p, q, r := 0, 1, 2
	for i := 3; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r

}

/**
 * @desc: 给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。
 * 一旦支付此费用，即可选择向上爬一个或者两个台阶。
 * 状态方程：dp[i]=min(dp[i−1]+cost[i−1],dp[i−2]+cost[i−2])
 * @data: 2022.1.23 12:08
 */
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = 0

	//dp[i]=min(dp[i−1]+cost[i−1],dp[i−2]+cost[i−2])
	for i := 2; i <= len(cost); i++ {
		//表示从下标 i-1 使用 cost[i−1] 的花费达到下标 i，或者从下标 i−2 使用 cost[i−2] 的花费达到下标 i
		dp[i] = getMin(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[len(cost)]
}

func getMin(i int, i2 int) int {
	if i < i2 {
		return i
	} else {
		return i2
	}
}
