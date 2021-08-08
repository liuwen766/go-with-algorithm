package selectsort

func Sort(arr []int) {
	//遍历，找到最小值，排到最前
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				SwapGo(arr, i, j)
			}
		}
	}
}
func SwapGo(list []int, i int, j int) {
	tmp := list[i]
	list[i] = list[j]
	list[j] = tmp
}
