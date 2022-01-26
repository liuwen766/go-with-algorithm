package main

import "fmt"

/*第一个，最基本的二分查找算法：
因为我们初始化 right = nums.length - 1
所以决定了我们的「搜索区间」是 [left, right]
所以决定了 while (left <= right)
同时也决定了 left = mid+1 和 right = mid-1

因为我们只需找到一个 target 的索引即可
所以当 nums[mid] == target 时可以立即返回
*/

/*第二个，寻找左侧边界的二分查找：
因为我们初始化 right = nums.length
所以决定了我们的「搜索区间」是 [left, right)
所以决定了 while (left < right)
同时也决定了 left = mid + 1 和 right = mid

因为我们需找到 target 的最左侧索引
所以当 nums[mid] == target 时不要立即返回
而要收紧右侧边界以锁定左侧边界
*/

/*第三个，寻找右侧边界的二分查找：
因为我们初始化 right = nums.length
所以决定了我们的「搜索区间」是 [left, right)
所以决定了 while (left < right)
同时也决定了 left = mid + 1 和 right = mid

因为我们需找到 target 的最右侧索引
所以当 nums[mid] == target 时不要立即返回
而要收紧左侧边界以锁定右侧边界

又因为收紧左侧边界时必须 left = mid + 1
所以最后无论返回 left 还是 right，必须减一
*/
//【这里统一以上三种搜索方法】
/**
 * @desc: 二分搜索的细节问题在于理解概念——「搜索区间」
 * @data: 2022.1.26 19:52
 */
func main() {
	nums := []int{-1, 0, 5, 5, 5, 5, 5, 9, 12}
	target := 5

	fmt.Println("寻找一个数的二分搜索：", search(nums, target))

	fmt.Println("寻找左侧边界的二分搜索：", searchLeft(nums, target))

	fmt.Println("寻找右侧边界的二分搜索：", searchRight(nums, target))
}

func searchRight(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}

func searchLeft(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

/**
 * @desc: 寻找一个数的二分搜索
 * @data: 2022.1.26 19:43
 */
func search(nums []int, target int) int {
	//不要出现 else，而是把所有情况用 else if 写清楚，这样可以清楚地展现所有细节
	left := 0
	right := len(nums) - 1
	//1、为什么 while 循环的条件中是 <=，而不是 <？
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
			//2、为什么 left = mid + 1，right = mid - 1？我看有的代码是 right = mid 或者 left = mid，没有这些加加减减，到底怎么回事，怎么判断？
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}
