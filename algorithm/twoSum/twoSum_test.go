package main

import (
	"testing"
)

func Test_twoSum(t *testing.T) {
	two1 := twoSum1([]int{2, 4, 5, 7}, 9)
	t.Logf("twoSum1:%v", two1)
	two2 := twoSum2([]int{2, 4, 5, 7}, 9)
	t.Logf("twoSum2:%v", two2)
}
