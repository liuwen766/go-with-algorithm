package insertsort

func Sort(list []int) {

	for i := 1; i < len(list); i++ { //从第二个数开始
		currentVal := list[i] //当前数
		insertIndex := i - 1  //下标
		//eg: 1 2 3 5 6 8 9 7
		for insertIndex >= 0 && list[insertIndex] > currentVal {
			list[insertIndex+1] = list[insertIndex] //后移
			insertIndex--
		}
		//完成插入
		if insertIndex+1 != i {
			list[insertIndex+1] = currentVal
		}
	}
}
