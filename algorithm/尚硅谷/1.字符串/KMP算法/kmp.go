package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "BBC ABCDAB ABCDABCDABDE"
	str2 := "ABCDABD"

	fmt.Println("工具类匹配结果:", stringsMatch(str1, str2))

	fmt.Println("暴力匹配结果:", violenceMatch(str1, str2))

	fmt.Println("获取部分匹配表:", getKmpNext(str2))
	fmt.Println("KMP结果:", kmpMatch(str1, str2))
}

/**
 * @desc: KMP算法
 * 移动位数=已匹配字符数-对应的部分匹配值
 * 部分匹配值，即 前缀 和 后缀 最长的共有元素的长度
 * @data: 2022.1.22 20:40
 */
func kmpMatch(str1 string, str2 string) int {
	// 1.得出部分匹配表 kmpNext数组
	next := getKmpNext(str2)
	// 2.计算移动位数
	for i, j := 0, 0; i < len(str1); i++ {

		for j > 0 && str1[i] != str2[j] {
			j = next[j-1]
		}

		if str1[i] == str2[j] {
			j++
		}
		if j == len(str2) {
			return i - j + 1
		}
	}
	return -1
}

/**
 * @desc: 得出部分匹配表 kmpNext数组
 * @data: 2022.1.23 11:34
 */
func getKmpNext(str string) []int {
	if len(str) == 0 {
		return []int{}
	}

	//创建一个next 数组保存部分匹配值
	next := make([]int, len(str))
	next[0] = 0

	for i, j := 1, 0; i < len(str); i++ {
		//当dest.charAt(i) != dest.charAt(j) ，我们需要从next[j-1]获取新的j
		//直到我们发现 有  dest.charAt(i) == dest.charAt(j)成立才退出
		//这时kmp算法的核心点
		for j > 0 && str[i] != str[j] {
			j = next[j-1]
		}

		//当dest.charAt(i) == dest.charAt(j) 满足时，部分匹配值就是+1
		if str[i] == str[j] {
			j++
		}
		next[i] = j
	}

	return next
}

/**
 * @desc: 暴力匹配
 * 移动位数=1
 * @data: 2022.1.22 20:38
 */
func violenceMatch(str1 string, str2 string) int {
	s1 := []rune(str1)
	s2 := []rune(str2)

	i := 0
	j := 0

	for i < len(s1) && j < len(s2) {
		if s1[i] == s2[j] {
			i++
			j++
		} else {
			i = i - (j - 1)
			j = 0
		}
	}

	if j == len(s2) {
		return i - j
	}
	return -1
}

/**
 * @desc: 暴力匹配
 * @data: 2022.1.22 20:38
 */
func stringsMatch(str1 string, str2 string) int {
	return strings.Index(str1, str2)
}
