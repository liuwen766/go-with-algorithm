package main

import (
	"fmt"
	"strings"
)

/**
 * @desc: 给你一个字符串 s，找到 s 中最长的回文子串。
 * @data: 2022.1.26 14:25
 */
func main() {
	s := "cbbd"
	fmt.Println("最长回文子串:", longestPalindrome(s))
	fmt.Println("暴力搜索获取最长回文子串:", longestPalindrome1(s))
}

/**
 * @desc: 从中心向两边扩展
 * @data: 2022.1.26 14:28
 */
func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		//以 s[i] 为中心的最长回文子串
		s1 := getPalindrome(s, i, i)
		//以 s[i] 和 s[i+1] 为中心的最长回文子串
		s2 := getPalindrome(s, i, i+1)
		res = getMaxLenString(res, s1, s2)
	}
	return res
}

func getMaxLenString(res string, s1 string, s2 string) string {
	ans := ""
	if len(res) > len(s1) {
		ans = res
	} else {
		ans = s1
	}
	if len(ans) < len(s2) {
		ans = s2
	}
	return ans
}

func getPalindrome(s string, l int, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}

func longestPalindrome1(s string) string {
	maxIndex := 0
	maxLen := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if isRevserse(s[i:j]) && len(s[i:j]) > maxLen {
				maxIndex = i
				maxLen = len(s[i:j])
			}
		}
	}
	return s[maxIndex : maxIndex+maxLen]
}

/**
 * @desc: 是否是翻转字符串
 * @data: 2022.1.22 16:27
 */
func isRevserse(s string) bool {
	runes := []rune(s)
	for left, right := 0, len(runes)-1; left < right; left, right = left+1, right-1 {
		runes[left], runes[right] = runes[right], runes[left]
	}
	return strings.EqualFold(s, string(runes))
}
