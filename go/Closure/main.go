package main

import "fmt"

func main() {

	fmt.Println("闭包就是一个函数和与其相关的引用环境组合的一个整体（实体）")
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))

	fmt.Println("文件后缀名处理案例。闭包的好处:如果使用传统的方法，也可以轻松实现这个功能，但是传统方法需要每次都传入后缀名，\n比如jpg,而闭包因为可以保留上次引用的某个值，所以我们传入一次就可以反复使用。")

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
