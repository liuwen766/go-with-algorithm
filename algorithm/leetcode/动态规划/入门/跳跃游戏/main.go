package main

import "fmt"

/**
 * @desc: 动态规划 跳跃游戏I 跳跃游戏II
 *        贪心算法 跳跃游戏II
 *  贪心选择性质，我们不需要「递归地」计算出所有选择的具体结果然后比较求最值，而只需要做出那个最有「潜力」，看起来最优的选择即可。
 * @data: 2022.1.25 11:08
 */
func main() {
	nums := []int{2, 3, 1, 1, 4}
	bool := canJump(nums)
	fmt.Println("是否能跳到最后一步：", bool)

	res := jump(nums)
	fmt.Println("跳到最后一步最少需要多少步：", res)
}

/**
 * @desc: 跳跃游戏I
 * 给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
 * 数组中的每个元素代表你在该位置可以跳跃的最大长度。
 * 判断你是否能够到达最后一个下标。
 * @data: 2022.1.25 11:16
 */
func canJump(nums []int) bool {
	jumpIndex := 0
	for i := 0; i < len(nums)-1; i++ {
		jumpIndex = getMax(jumpIndex, i+nums[i])
		// 可能碰到 nums[i]=0，卡住跳不动
		if jumpIndex <= i {
			return false
		}
	}
	if jumpIndex >= len(nums)-1 {
		return true
	}
	return false
}

/**
 * @desc: 跳跃游戏II
给你一个非负整数数组nums ，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
你的目标是使用最少的跳跃次数到达数组的最后一个位置。
假设你总是可以到达数组的最后一个位置。
 * @data: `2022.1.25` 11:15
*/
func jump(nums []int) int {
	jump := 0
	end := 0
	farthest := 0
	for i := 0; i < len(nums)-1; i++ {
		farthest = getMax(farthest, i+nums[i])
		if end == i {
			jump++
			end = farthest
		}
	}

	return jump
}

func getMax(index int, i int) int {
	if index > i {
		return index
	}
	return i
}
