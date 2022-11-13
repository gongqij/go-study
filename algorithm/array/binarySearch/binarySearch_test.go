package search

import "testing"

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	target := 5
	targetIndex := binarySearch(nums, target)
	t.Logf("数组:%v,target:%d,targetIndex:%d\n", nums, target, targetIndex)
}
