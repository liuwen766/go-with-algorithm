package main

import "fmt"

func main() {
	coins := []int{1, 2, 5, 10}
	amount := 50
	change1 := coinChange1(coins, amount)
	change2 := coinChange2(coins, amount)
	change3 := coinChange3(coins, amount)

	fmt.Println(change1, change2, change3)

}

/**
 * @desc: 暴力递归【先自顶向下，再自底向上】
 * @data: 2022/12/10 14:23
 */
func coinChange1(coins []int, amount int) int {
	return 0
}

/**
 * @desc: 带备忘录的递归
 * @data: 2022/12/10 14:23
 */
func coinChange2(coins []int, amount int) int {
	return 0
}

/**
 * @desc: dp 数组的迭代解法【自底向上使用 dp table 来消除重叠子问题】
 * @data: 2022/12/10 14:23
 */
func coinChange3(coins []int, amount int) int {
	return 0
}
