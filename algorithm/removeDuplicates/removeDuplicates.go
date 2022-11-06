package removeDuplicates

//升序数组去重
func ascArrayRemoveDuplication(nums []int) []int {
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return nums[:slow]
}

//无序数组去重1
func arrayRemoveDuplication1(nums []int) []int {
	temp := make(map[int]struct{})
	for _, num := range nums {
		temp[num] = struct{}{}
	}
	var res []int
	for k := range temp {
		res = append(res, k)
	}
	return res
}

//无序数组去重2
func arrayRemoveDuplication2(nums []int) []int {
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
