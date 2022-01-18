package heapsort

func Sort(arr []int) {
	//step1：构造大顶堆
	for i := len(arr)/2 - 1; i >= 0; i-- {
		adjust(arr, i, len(arr))
	}
	//fmt.Println("arr:", arr)
	//step2：堆顶元素交换堆尾元素，重新调整
	for i := len(arr) - 1; i >= 0; i-- {
		swap(arr, 0, i)
		adjust(arr, 0, i)
	}
}

func adjust(arr []int, i int, len int) {
	tmp := arr[i]
	for k := 2*i + 1; k < len; k = 2*k + 1 {
		if k+1 < len && arr[k] < arr[k+1] {
			k++
		}
		if arr[k] > tmp {
			arr[i] = arr[k]
			i = k
		} else {
			break
		}
	}
	arr[i] = tmp
}

func swap(arr []int, i int, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
