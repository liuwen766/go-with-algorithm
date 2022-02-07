package main

import "fmt"

/**
 * @desc: 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案
 * @data: 2022.1.27 23:10
 */
func main() {
	nums := []int{5, 6, 4, 2}
	fmt.Println("数组", nums, "的全排列是：", permute(nums))
}

/*
解决一个回溯问题，实际上就是一个决策树的遍历过程。只需要思考 3 个问题：
1、路径：也就是已经做出的选择。
2、选择列表：也就是你当前可以做的选择。
3、结束条件：也就是到达决策树底层，无法再做选择的条件。
*/
var res [][]int
var visited map[int]bool

func permute(nums []int) [][]int {
	res = make([][]int, 0)
	visited = make(map[int]bool)
	track := make([]int, 0)
	backTrack(nums, track)
	return res
}

func backTrack(nums []int, track []int) {
	// 触发结束条件
	if len(track) == len(nums) {
		//加入解集时，要将数组内容拷贝到一个新的数组里，再加入解集
		temp := make([]int, len(track))
		copy(temp, track)
		res = append(res, temp)
		// 触发结束条件
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		visited[i] = true
		track = append(track, nums[i])
		backTrack(nums, track)
		track = track[:len(track)-1]
		visited[i] = false
	}
}
