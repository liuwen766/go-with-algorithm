package main

import "fmt"

func main() {
	var nums = []int{3, 3, 1}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	var res int
	for _, val := range nums {
		res = res ^ val
	}
	return res
}
