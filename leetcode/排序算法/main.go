package main

import (
	"fmt"
	"github.com/go-with-algorithm/leetcode/排序算法/bubblesort"
	"github.com/go-with-algorithm/leetcode/排序算法/bucketsort"
	"github.com/go-with-algorithm/leetcode/排序算法/countsort"
	"github.com/go-with-algorithm/leetcode/排序算法/heapsort"
	"github.com/go-with-algorithm/leetcode/排序算法/insertsort"
	"github.com/go-with-algorithm/leetcode/排序算法/mergesort"
	"github.com/go-with-algorithm/leetcode/排序算法/quicksort"
	"github.com/go-with-algorithm/leetcode/排序算法/radixsort"
	"github.com/go-with-algorithm/leetcode/排序算法/selectsort"
	"github.com/go-with-algorithm/leetcode/排序算法/shellsort"
)

var data = []int{8, 3, 6, 9, 11, 2, 7, 23, 65, 13, 9}

func main() {
	//datePrintln("桶排序")
	//datePrintln("计数排序")
	datePrintln("冒泡排序")
	//datePrintln("快速排序")
	//datePrintln("选择排序")
	//datePrintln("插入排序")
	//datePrintln("希尔排序")
	//datePrintln("合并排序")
	//datePrintln("基数排序")
	//datePrintln("堆排序")
}

func datePrintln(name string) {
	var result []int
	fmt.Println("初始数据:", data)
	switch name {
	case "冒泡排序":
		//bubblesort.Sort(data, 0, len(data)-1)
		bubblesort.MyBubbleSort(data)
		result = data
		break
	case "快速排序":
		quicksort.Sort(data, 0, len(data)-1)
		result = data
		break
	case "选择排序":
		selectsort.Sort(data)
		result = data
		break
	case "插入排序":
		insertsort.Sort(data, 0, len(data)-1)
		result = data
		break
	case "合并排序":
		result = mergesort.Sort(data, 0, len(data)-1)
		break
	case "希尔排序":
		shellsort.Sort(data, 0, len(data)-1)
		result = data
		break
	case "堆排序":
		heapsort.Sort(data)
		result = data
		break
	case "基数排序":
		radixsort.Sort(data)
		result = data
		break
	case "桶排序":
		result = bucketsort.Sort(data)
		break
	case "计数排序":
		result = countsort.Sort(data)
		break
	}
	fmt.Println(name+":", result)
}
