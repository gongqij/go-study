package removeElement

//移除元素双指针
func removeElement(nums []int, val int) []int {
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] != val {
			nums[left] = nums[right]
			left++
		}
	}
	return nums[:left]
}

//移除元素:双指针优化
//避免了需要保留的元素的重复赋值操作
func removeElement2(nums []int, val int) []int {
	left, right := 0, len(nums)
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return nums[:left]
}
