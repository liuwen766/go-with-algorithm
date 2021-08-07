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
	//遍历，找到最小值，排到最前
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				SwapGo(arr, i, j)
			}
		}
	}
}
