package main

import "testing"

func Test_climbStairs(t *testing.T) {
	for i := 1; i <= 10; i++ {
		t.Logf("爬%d级台阶共有%d种方法", i, climbStairs2(0))
	}
}
