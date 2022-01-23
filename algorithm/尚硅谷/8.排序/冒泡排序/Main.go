package main

import "fmt"

func main() {

	arr := []int{2, 5, 3, 9, 7, 6, 1, 4, 8}
	bubbleSort(arr)
	fmt.Println("冒泡排序后：", arr)
}

func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				swap(arr, i, j)
			}
		}
	}
}

func swap(arr []int, i int, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
