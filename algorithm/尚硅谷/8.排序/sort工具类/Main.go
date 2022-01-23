package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{2, 5, 3, 9, 7, 6, 1, 4, 8}
	fmt.Println("未排序前，sort.IntsAreSorted", sort.IntsAreSorted(arr))
	sort.Ints(arr)
	sort.SearchInts(arr, 3)
	fmt.Println("sort.Ints()工具类：", arr)
	fmt.Println("排序后，sort.IntsAreSorted", sort.IntsAreSorted(arr))

	str := []string{"r", "c", "d", "a", "f", "b"}
	sort.Strings(str)
	sort.StringsAreSorted(str)
	fmt.Println("sort.Strings()工具类：", str)

	sort.Float64s(nil)
	sort.Sort(nil)
	sort.IsSorted(nil)
	sort.Stable(nil)
}
