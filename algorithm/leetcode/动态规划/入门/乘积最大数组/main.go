package main

import (
	"fmt"
	"math"
)

/**
 * @desc: 给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
 * @data: 2022.1.27 9:28
 */
func main() {
	nums := []int{0, 2}
	fmt.Println("乘积最大子数组对应的乘积为：", maxProduct(nums))
}

func maxProduct(nums []int) int {
	max := math.MinInt8
	return max
}
