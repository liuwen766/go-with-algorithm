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
	need := make(map[uint8]int, 0)
	window := make(map[uint8]int, 0)
	for i := range s1 {
		need[s1[i]]++
	}
	left, right := 0, 0
	valid := 0
	for right < len(s2) {
		c := s2[right]
		right++
		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		// 判断左侧窗口是否要收缩
		for right-left >= len(s1) {
			// 找到符合条件的子串
			if valid == len(need) {
				return true
			}
			d := s2[left]
			left++
			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	// 未找到符合条件的子串
	return false
}
