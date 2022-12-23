package main

import "strings"

func main() {

}

//验证回文串：https://leetcode.cn/problems/valid-palindrome/
//如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
//给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。
func isPalindrome(s string) bool {
	lower := strings.ToLower(s)
	left, right := 0, len(lower)-1
	for left < right {
		//	从左向右找出第一个字符
		for left < right && isValid(lower[left]) {
			left++
		}
		//	从右向左找出第一个字符
		for left < right && isValid(lower[right]) {
			right--
		}
		//对比，如果不一致，返回false
		if left < right {
			if lower[left] != lower[right] {
				return false
			}
			left--
			right++
		}
	}
	return true
}

func isValid(c byte) bool {
	if (c > 'a' && c < 'z') || (c > '0' && c < '9') {
		return true
	}
	return false
}
