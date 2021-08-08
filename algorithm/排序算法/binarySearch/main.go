package main

import "fmt"

func main() {
	//有序数组+目标数据
	arr := []int{1, 5, 9, 12, 26, 89, 156, 164, 962}
	target := 12
	//res := binarySearch1(arr, 0, len(arr)-1, target)
	res := binarySearch2(arr, 0, len(arr)-1, target)
	if res != -1 {
		fmt.Printf("target:%d在index:%d的数组上,数组arr[%d]=%d", target, res, res, arr[res])
	} else {
		fmt.Printf("target:%d不在数组arr中", target)
	}
}

/**
 * @desc: 非递归
 * @data: 2021.8.7 16:52
 */
func binarySearch1(arr []int, left, right, target int) int {
	if len(arr) == 0 {
		return -1
	}
	left = 0
	right = len(arr) - 1
	var mid int
	for left <= right {
		mid = (left + right) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

/**
 * @desc: 递归方式
 * @data: 2021.8.7 16:54
 */
func binarySearch2(arr []int, left, right, target int) int {
	if len(arr) == 0 || left > right {
		return -1
	}
	var res int
	if arr[(left+right)/2] == target {
		res = (left + right) / 2
	} else if arr[(left+right)/2] > target {
		res = binarySearch2(arr, left, (left+right)/2-1, target)
	} else {
		res = binarySearch2(arr, (left+right)/2+1, right, target)
	}
	return res
}
