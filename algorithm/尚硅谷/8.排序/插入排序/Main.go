package main

import (
	"fmt"
)

func main() {

	arr := []int{4, 3, 2, 10, 12, 1, 5, 6, 8}
	insertSort(arr)
	fmt.Println("插入排序后：", arr)
}

func insertSort(arr []int) {
	//从第1个数开始
	for i := 1; i < len(arr); i++ {
		//拿出第0个数
		j := i - 1
		//从后往前  两两比较  比较完之后，完成第一遍插入排序
		for j > 0 && arr[j] < arr[j-1] {
			swap(arr, j, j-1)
			j--
		}
	}
}

func swap(arr []int, i int, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
