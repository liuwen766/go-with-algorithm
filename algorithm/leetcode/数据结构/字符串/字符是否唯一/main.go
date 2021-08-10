package main

import "fmt"

func main() {
	fmt.Println(isUnique("leetcode"))
}
func isUnique(astr string) bool {

	var hashStr = make(map[byte]struct{})
	for i := 0; i < len(astr); i++ {
		if _, ok := hashStr[astr[i]]; ok {
			return false
		}
		hashStr[astr[i]] = struct{}{}
	}
	return true
}
