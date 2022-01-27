package main

import "fmt"

/**
 * @desc: 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
 * @data: 2022.1.27 13:39
 */
func main() {
	s := "abcabcbb"
	fmt.Println("最长无重复子串的长度：", lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {

	maxLen := 0
	return maxLen
}

func getMax(maxLen int, i int) int {
	if maxLen > i {
		return maxLen
	}
	return i
}
