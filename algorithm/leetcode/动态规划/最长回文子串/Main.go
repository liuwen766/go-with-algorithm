package main

import (
	"fmt"
	"strings"
)

func main() {
	res := longestPalindrome("aka")
	fmt.Println("最长回文子串:", res)

	fmt.Println("是否是回文子串:", isRevserse("abcba"))
}

/**
 * @desc: 动态规划
 * 【状态方程：P(i,j)=P(i+1,j−1)∧(Si==Sj)  即只有s[i+1:j-1]s[i+1:j−1] 是回文串，并且s的第i和j个字母相同时，s[i:j]才会是回文串。
 *   边界条件：P(i,i)=true
 *           P(i,i+1)=(Si==Si+1)】
 * @data: 2022.1.22 17:03
 */
func longestPalindrome(s string) interface{} {
	return ""
}

/**
 * @desc: 暴力求解【费时】
 * @data: 2022.1.22 16:13
 */
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
