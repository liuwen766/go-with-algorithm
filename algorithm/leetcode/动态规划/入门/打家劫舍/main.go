package main

import "fmt"

/**
 * @desc: 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，
影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房
屋在同一晚上被小偷闯入，系统会自动报警。给定一个代表每个房屋存放金额的非负整数
数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
 * @data: 2022.1.24 9:18
*/
func main() {
	nums := []int{2, 3, 2}
	fmt.Println("最高可以偷窃到:", rob(nums))

	fmt.Println("房屋是一个圈的情况下,最高可以偷窃到:", rob2(nums))
}

/**
 * @desc: 房屋是一个圈的情况【传参打劫范围】【注意边界问题】
 * @data: 2022.1.24 17:19
 */
func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	res := getMax(robRange(nums, 0, len(nums)-2), robRange(nums, 1, len(nums)-1))
	return res
}

func robRange(nums []int, start int, end int) int {
	dp_i_1 := 0
	dp_i_2 := 0
	dp_i := 0
	for i := end; i >= start; i-- {
		dp_i = getMax(dp_i_1, dp_i_2+nums[i])
		dp_i_2 = dp_i_1
		dp_i_1 = dp_i
	}
	return dp_i
}

/**
 * @desc: 状态方程：dp[n]=max[dp[n-1],dp[n-2]+nums[n]]
 * dp[0]=nums[0]
 * dp[1]=max[nums[0],nums[1]]
 * dp[2]=max[dp[1],dp[0]+nums[2]]
 * dp[3]=max[dp[2],dp[1]+nums[3]]
 * @data: 2022.1.24 9:23
 */
func rob(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	if len(nums) == 1 {
		return dp[0]
	}
	dp[1] = getMax(nums[0], nums[1])
	//状态方程：dp[n]=max[dp[n-1],dp[n-2]+nums[n]]
	for i := 2; i < len(nums); i++ {
		dp[i] = getMax(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]
}

func getMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
