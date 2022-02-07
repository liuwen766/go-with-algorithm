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
	window := make(map[uint8]int, 0)
	left, right := 0, 0
	for right < len(s) {
		c := s[right]
		right++
		// 进行窗口内数据的一系列更新
		window[c]++
		for window[c] > 1 {
			d := s[left]
			left++
			window[d]--
		}
		// 在这里更新答案
		maxLen = getMax(maxLen, right-left)
	}
	return maxLen
}

func getMax(maxLen int, i int) int {
	if maxLen > i {
		return maxLen
	}
	return i
}
