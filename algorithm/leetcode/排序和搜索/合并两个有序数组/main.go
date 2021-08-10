package main

import "fmt"

func main() {
	nums1 := []int{1, 3, 7, 0, 0, 0}
	nums2 := []int{2, 4, 6}
	merge(nums1, len(nums1)-3, nums2, len(nums2))
	fmt.Println(nums1)
}

/**
 * @desc: m 和 n 是数组的容量
 * @data: 2021.8.10 22:40
 */
func merge(nums1 []int, m int, nums2 []int, n int) {
	nums := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for {
		if p1 == m {
			nums = append(nums, nums2[p2:]...)
			break
		}
		if p2 == n {
			nums = append(nums, nums1[p1:]...)
			break
		}
		if nums1[p1] < nums2[p2] {
			nums = append(nums, nums1[p1])
			p1++
		} else {
			nums = append(nums, nums2[p2])
			p2++
		}
	}
	copy(nums1, nums)
}
