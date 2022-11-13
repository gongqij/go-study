package removeDuplication

//有序数组去重
func orderedArrayRemoveDuplicates(nums []int) []int {
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return nums[:slow]
}

//无序数组去重
func unorderedArrayRemoveDuplication(nums []int) []int {
	temp := make(map[int]struct{})
	var res []int
	for _, num := range nums {
		if _, ok := temp[num]; !ok {
			res = append(res, num)
		}
		temp[num] = struct{}{}
	}
	return res
}

//无序数组去重2
func unorderedArrayRemoveDuplication2(nums []int) []int {
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		flag := false
		for j := fast - 1; j >= 0; j-- {
			if nums[fast] == nums[j] {
				flag = true
				break
			}
		}
		if !flag {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return nums[:slow]
}
