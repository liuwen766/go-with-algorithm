package main

import (
	"fmt"
)

func main() {
	s := "We are first"
	space := replaceSpace(s)
	fmt.Println(space)
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * @param s string字符串
 * @return string字符串
 */
func replaceSpace(s string) string {

	var newStr []byte
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			newStr = append(newStr, '%')
			newStr = append(newStr, '2')
			newStr = append(newStr, '0')
		} else {
			newStr = append(newStr, s[i])
		}
	}
	return string(newStr)
}
