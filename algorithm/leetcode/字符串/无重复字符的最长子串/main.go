package main

import "fmt"

func main() {

	fmt.Println("不含有重复字符的最长子串的长度:", lengthOfLongestSubstring("pwwkew"))

}

/**
 * @desc: 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
 * 滑动窗口代码框架【双指针】
int left = 0, right = 0;
while (right < s.size()) {
    // 增大窗口
    window.add(s[right]);
    right++;
    while (window needs shrink) {
        // 缩小窗口
        window.remove(s[left]);
        left++;
    }
}
 * @data: 2022.1.23 16:40
*/
func lengthOfLongestSubstring(s string) int {
	res := 0
	left := 0
	right := 0

	return getMax(res, right-left)
}

func getMax(res int, i int) int {
	return 0
}
