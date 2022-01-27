package main

import "fmt"

/**
 * @desc: 给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
 * @data: 2022.1.27 11:07
 */
func main() {
	s1 := "hello"
	s2 := "ooolleoooleh"
	fmt.Println("字符串s2是否包含字符串s1的排列", checkInclusion(s1, s2))
}

func checkInclusion(s1 string, s2 string) bool {

	return false
}
