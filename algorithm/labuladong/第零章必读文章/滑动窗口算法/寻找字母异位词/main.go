package main

import "fmt"

/**
 * @desc: 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
 * @data: 2022.1.27 15:17
 */
func main() {
	s := "cbaebacbacd"
	p := "abc"
	fmt.Println("找出所有的异位词：", findAnagrams(s, p))
}

func findAnagrams(s string, p string) []int {

	ans := make([]int, 0)

	return ans
}
