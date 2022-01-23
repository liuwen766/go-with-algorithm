package main

import "fmt"

func main() {
	fmt.Println(fib(50))
}

/**
 * @desc: 滚动数组思想
 * @data: 2021.8.9 23:06
 */
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	p, q, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

//func fib(n int) int {
//	if n == 0 || n == 1 {
//		return n
//	}
//	return fib(n-1) + fib(n-2)
//}
