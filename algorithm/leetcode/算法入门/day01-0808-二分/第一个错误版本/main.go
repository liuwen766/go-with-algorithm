package main

func main() {
	firstBadVersion(1)
}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	if n < 1 {
		return 0
	}
	left := 1
	right := n
	for left <= right {
		mid := (left + right) / 2
		if isBadVersion(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func isBadVersion(version int) bool {
	return false
}
