package main

import "fmt"

/**
 * @desc: 给你一个整数数组nums，你可以对它进行一些操作。
每次操作中，选择任意一个nums[i]，删除它并获得nums[i]的点数。
之后，你必须删除 所有 等于nums[i] - 1 和 nums[i] + 1的元素。
开始你拥有 0 个点数。返回你能通过这些操作获得的最大点数。
 * @data: 2022.1.24 22:09
*/
func main() {
	nums := []int{2, 2, 3, 3, 3, 4}
	fmt.Println("获得的最大点数等价于能抢到的最多钱数:", deleteAndEarn(nums))
}

/**
 * @desc: 即打家劫舍问题：dp[i] = Math.max(dp[i - 1], dp[i - 2] + sum[i])
 * nums = [2, 2, 3, 3, 3, 4]
 * sums = [0, 0, 4, 9, 4];
 * 为啥说是打家劫舍问题？
 * 不能同时选择邻居 = 同步删除的 -1 AND +1 元素（对应数组 index）
 * 可以抢劫的金额 = 删除这个元素能获得的点数
 * @data: 2022.1.24 22:31
 */
func deleteAndEarn(nums []int) int {
	maxVal := 0
	for _, val := range nums {
		maxVal = max(maxVal, val)
	}
	sum := make([]int, maxVal+1)
	for _, val := range nums {
		sum[val] += val
	}
	return rob(sum)
}

func rob(sums []int) int {

	dp := make([]int, len(sums))
	dp[0] = sums[0]
	dp[1] = max(sums[0], sums[1])

	//dp[i] = Math.max(dp[i - 1], dp[i - 2] + sum[i])
	for i := 2; i < len(sums); i++ {
		dp[i] = max(dp[i-2]+sums[i], dp[i-1])
	}
	return dp[len(sums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
