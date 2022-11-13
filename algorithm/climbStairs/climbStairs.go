package main

//爬梯子==斐波那契数列
//方法一:动态规划-滚动数组
func climbStairs1(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

//方法二:递归
func climbStairs2(n int) int {
	if n < 2 {
		return 1
	}
	return climbStairs2(n-1) + climbStairs2(n-2)
}
