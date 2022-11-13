package main

import (
	"sort"
)

//直接合并后排序
func mergeOrderedArray1(nums1 []int, nums2 []int) []int {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	return nums1
}

//双指针
//时间复杂度：O(m+n)
//空间复杂度：O(m+n)
func mergeOrderedArray2(left []int, right []int) []int {
	var result []int
	for {
		if len(left) == 0 {
			result = append(result, right...)
			break
		}
		if len(right) == 0 {
			result = append(result, left...)
			break
		}
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	return result
}

//逆向双指针
//时间复杂度：O(m+n)
//空间复杂度：O(m+n)
func mergeOrderedArray3(nums1 []int, nums2 []int) []int {
	m := len(nums1)
	n := len(nums2)
	nums1 = append(nums1, nums2...)
	for p1, p2, tail := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; tail-- {
		var cur int
		if p1 == -1 {
			cur = nums2[p2]
			p2--
		} else if p2 == -1 {
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur
	}
	return nums1
}
