package 查找元素的第一个和最后一个位置

/**
 * @desc: 给定一个n个元素有序的（升序）整型数组nums和一个目标值target，写一个函数搜索ums中的target,
 * 如果目标值存在返回下标，否则返回-1。
 * @data: 2022/12/10 15:31
 */
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		// ***
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			return mid
		}
	}
	return -1
}

/**
 * @desc: 给你一个按照非递减顺序排列的整数数组ms，和一个目标值target。请你找出给定目标值在数组中的开始位置和结束位置。
 * 如果数组中不存在目标值target,返回[-1，-1]。
 * @data: 2022/12/10 16:08
 */
func searchRange(nums []int, target int) []int {
	return []int{leftBondSearch(nums, target), rightBondSearch(nums, target)}
}

// 找到最左边的值
func leftBondSearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			// 缩小右边界【注意】
			right = mid - 1
		}
	}
	// [2,2,2,2,2,2,2] 3
	// 说明没有找到【注意】
	if left == len(nums) {
		return -1
	}
	// 判断是不是找到
	if nums[left] == target {
		return left
	}
	return -1
}

// 找到最右边的值
func rightBondSearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			// 缩小左边界【注意】
			left = mid + 1
		}
	}
	// 说明没有找到【注意】
	if left-1 < 0 {
		return -1
	}
	// 判断是不是找到
	if nums[left-1] == target {
		return left - 1
	}
	return -1
}
