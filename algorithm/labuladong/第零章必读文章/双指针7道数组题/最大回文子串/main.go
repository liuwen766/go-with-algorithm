package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("babad"))
}

//最长回文子串【核心必备】
func longestPalindrome(s string) string {
	res := ""
	// 双指针  从中心向两边扩散
	for i := 0; i < len(s); i++ {
		// 获取以i,i为中心的最长回文串【长度为奇数】
		s1 := isPalindrome(s, i, i)
		// 获取以i,i+1为中心的最长回文串【长度为偶数】
		s2 := isPalindrome(s, i, i+1)
		// 获取最大回文串
		res = getMaxInThree(res, s1, s2)
	}
	return res
}

func getMaxInThree(s1 string, s2 string, s3 string) string {
	if len(s1) > len(s2) {
		if len(s1) > len(s3) {
			return s1
		} else {
			return s3
		}
	} else {
		if len(s2) > len(s3) {
			return s2
		} else {
			return s3
		}
	}
}

func isPalindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}
