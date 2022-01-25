package main

import "fmt"

/**
 * @desc: 给你 k 种面值的硬币，面值分别为 c1, c2 ... ck，每种硬币的数量无限，
 * 再给一个总金额 amount，问你最少需要几枚硬币凑出这个金额，如果不可能凑出，算法返回 -1 。
 * @data: 2022.1.22 17:21
 */
func main() {
	res := coinChange([]int{1, 2, 5}, 58)
	fmt.Println("最少需要", res, "枚硬币！")
}

/**
 * @desc: 动态规划算法
 * 凑成0元需要0个硬币   //d(0)=0
 * 凑成1元需要1个1元硬币 //d(1)=d(0)+1=1
 * 凑成2元需要2个1元硬币或1个2元硬币 那么我们选择1个2元硬币 //d(2)=min{d(1)+1,d(2-2)}=1
 * 凑成3元需要3个1元硬币或者1个1元硬币1个2元硬币，那么我们选择1个1元硬币1个2元硬币 //d(3)=min{d(2)+1}=2
 * ......
 * 抽离出来d(i)=min{ d(i-1)+1,d(i-vj)+1 }，其中i-vj >=0，vj表示第j个硬币的面值;
 * @data: 2022.1.22 19:17
 */
func coinChange(arr []int, target int) int {
	return dp(arr, target)
}

func dp(arr []int, target int) int {
	if target < 0 {
		return -1
	}
	if target < 2 {
		return target
	}
	for i := 0; i < target; i++ {

	}
	return -1
}
