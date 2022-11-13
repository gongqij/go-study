package search

// 有序数组的二分查找
// 二分法的关键思想是
// 假设该数组的长度是N那么二分后是N/2，再二分后是N/4……直到二分到1结束（当然这是属于最坏的情况了，即每次找到的那个中点数都不是我们要找的），
// 那么二分的次数就是基本语句执行的次数，于是我们可以设次数为x，N * (1/2)^x=1；则x=logN,底数是2，
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		num := nums[mid]
		if num == target {
			return mid
		} else if num > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
