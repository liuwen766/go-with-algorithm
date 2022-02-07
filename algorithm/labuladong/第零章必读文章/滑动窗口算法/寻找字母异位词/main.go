package main

import "fmt"

/**
 * @desc: 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
 * @data: 2022.1.27 15:17
 */
func main() {
	s := "baa"
	p := "aa"
	fmt.Println("找出所有的异位词：", findAnagrams(s, p))
}

func findAnagrams(s string, p string) []int {
	ans := make([]int, 0)
	need := make(map[uint8]int, 0)
	window := make(map[uint8]int, 0)
	for i := range p {
		need[p[i]]++
	}
	left, right := 0, 0
	valid := 0
	for right < len(s) {
		c := s[right]
		right++
		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for right-left >= len(p) {
			if valid == len(need) {
				ans = append(ans, left)
			}
			d := s[left]
			left++
			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}

	}
	return ans
}
