package main

import "fmt"

func main() {
	fmt.Println(tribonacci(25))
}

func tribonacci(n int) int {
	if n < 2 {
		return n
	}
	if n == 2 {
		return 1
	}
	o, p, q, r := 0, 0, 1, 1
	for i := 3; i <= n; i++ {
		o = p
		p = q
		q = r
		r = o + p + q
	}
	return r
}
