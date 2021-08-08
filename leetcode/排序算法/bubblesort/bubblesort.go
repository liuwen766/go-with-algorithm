package bubblesort

func Sort(list []int, left, right int) {
	if right == 0 {
		return
	}
	for index, num := range list {
		if index < right && num > list[index+1] {
			SwapGo(list, index, index+1)
		}
	}
	Sort(list, left, right-1)
}

func SwapGo(list []int, i int, j int) {
	tmp := list[i]
	list[i] = list[j]
	list[j] = tmp
}

func MyBubbleSort(arr []int) {
	//冒泡
	for i := len(arr); i > 0; i-- {
		for j := len(arr) - 1; j > 0; j-- {
			if arr[j-1] > arr[j] {
				SwapGo(arr, j-1, j)
			}
		}
	}
}
