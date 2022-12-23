package main

import "fmt"

func main() {
	res1 := fib1(10)
	res2 := fib2(10)
	res3 := fib3(10)
	res4 := fib4(10)
	fmt.Println(res1, res2, res3, res4)
}

// 暴力递归【先自顶向下，再自底向上】
func fib1(n int) int {
	return 0
}

// 带备忘录的递归【减少递归的大量重复计算】
func fib2(n int) int {
	return 0
}

// 自底向上求解【dp 数组的迭代解法1】
func fib3(n int) int {
	return 0
}

// 自底向上求解【dp 数组的迭代解法2】
func fib4(n int) int {
	return 0
}
