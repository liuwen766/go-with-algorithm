package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {

	usedRate := Decimal4(0.72324545, 4)
	fmt.Println("usedRate", usedRate)
	content := fmt.Sprintf("资源池 %v (%v) %v云盘的使用率已达%.2f%%\n", "suzhou", "可用区2", "G1", usedRate*100)
	fmt.Println("content", content)

	fmt.Println("闭包就是一个函数和与其相关的引用环境组合的一个整体（实体）")
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))

	fmt.Println("文件后缀名处理案例。闭包的好处:如果使用传统的方法，也可以轻松实现这个功能，但是传统方法需要每次都传入后缀名，\n" +
		"比如jpg,而闭包因为可以保留上次引用的某个值，所以我们传入一次就可以反复使用。")

	//闭包=函数+引用环境
	var g = adder()
	fmt.Println(g(10)) //10
	fmt.Println(g(20)) //30
	fmt.Println(g(30)) //60

	g1 := adder()
	fmt.Println(g1(40)) //40
	fmt.Println(g1(50)) //90

	err := fmt.Sprintf("Get metric [%s] from prometheus try 2 times still has no return! perhaps prometheus is down! please check! SendFlag is: %v", "query", false)
	err2 := errors.New(err)
	fmt.Println(err2)

}

func Decimal4(value float64, prec int) float64 {
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', prec, 64), 64)
	return value
}

/**
 * @desc: 累加器（闭包实现）
 * @data: 2021.8.7 13:06
 */
func AddUpper() func(int) int {
	n := 10
	str := "hello"
	return func(i int) int {
		n = n + i
		str += string(36)
		fmt.Println(str)
		return n
	}
}

func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}
