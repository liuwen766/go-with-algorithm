package main

import "fmt"

func main() {

	var nums = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))

}

/**
 * @desc: f(i)=max{f(i−1)+nums[i],nums[i]}
 * @data: 2021.8.9 21:22
 */
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}
