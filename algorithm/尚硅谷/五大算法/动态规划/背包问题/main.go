package main

import (
	"fmt"
)

func main() {
	weight := []int{1, 4, 3}         //物品的重量
	value := []int{1500, 3000, 2000} //物品的价值
	capacity := 4                    //背包的容量

	maxVal := BeiBao(weight, value, capacity)

	fmt.Println("最大价值为：", maxVal)

}

/**
 * @desc: 背包问题
 * 状态转移方程  dp[i][j] = max(dp[i−1][j], dp[i−1][j−w[i]]+v[i]) // j >= w[i]
 * @data: 2022.1.22 19:59
 */
func BeiBao(weight []int, value []int, capacity int) int {
	//dp[i][j]表示将前i件物品装进限重为j的背包可以获得的最大价值, 0<=i<=N, 0<=j<=W
	var v [][]int
	for i := 0; i < len(value); i++ {
		ar := make([]int, capacity)
		v = append(v, ar)
	}

	return 0
}
