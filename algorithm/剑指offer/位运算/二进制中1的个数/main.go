package main

import "fmt"

func main() {
	var num uint32
	num = 123
	weight := hammingWeight(num)
	fmt.Println(weight)
}

func hammingWeight(num uint32) int {
	res := uint32(0)
	for num != uint32(0) {
		res += num & 1
		num >>= 1
	}
	return int(res)
}
