package main

import "fmt"

func main() {
	arr := []int{4, 3, 2, 10, 12, 1, 5, 6, 8}
	selectSort(arr)
	fmt.Println("选择排序后：", arr)
}

func selectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		//假设最小值是第一个值
		tmp := arr[i]
		minIndex := i
		//找出真正的最小值
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		//遍历完一遍后，肯定拿到最小值
		arr[i] = arr[minIndex]
		arr[minIndex] = tmp
	}
}
