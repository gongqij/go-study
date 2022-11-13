package maxSubArray

//最大连续子序和

//动态规划
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1] //前边的元素大于0，加到当前元素上
		}
		//nums[i]保存着nums[i]+nums[i-1]+nums[i-2]...+nums[j]的连续子序和，nums[j]一定大于0,且[0,j)区间的连续子序和<=0
		//nums[i]有可能小于max
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// 贪心法
//TODO 未学习
func maxSubArray2(nums []int) int {
	n := len(nums)
	// 这里的dp[i] 表示，最大的连续子数组和，包含num[i] 元素
	dp := make([]int, n)
	// 初始化，由于dp 状态转移方程依赖dp[0]
	dp[0] = nums[0]
	// 初始化最大的和
	mx := nums[0]
	for i := 1; i < n; i++ {
		// 这里的状态转移方程就是：求最大和
		// 会面临2种情况，一个是带前面的和，一个是不带前面的和
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		mx = max(mx, dp[i])
	}
	return mx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
