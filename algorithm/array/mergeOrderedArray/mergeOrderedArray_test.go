package main

import (
	"testing"
)

func Test_MergeOrderedArray(t *testing.T) {
	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{3, 4, 5, 6, 7}

	t.Run("mergeOrderedArray1", func(t *testing.T) {
		t.Logf("before merge:arr1:%v,arr2:%v", arr1, arr2)
		merged := mergeOrderedArray1(arr1, arr2)
		t.Logf("after merge:%v", merged)
	})
	t.Run("mergeOrderedArray2", func(t *testing.T) {
		t.Logf("before merge:arr1:%v,arr2:%v", arr1, arr2)
		merged := mergeOrderedArray2(arr1, arr2)
		t.Logf("after merge:%v", merged)
	})
	t.Run("mergeOrderedArray3", func(t *testing.T) {
		t.Logf("before merge:arr1:%v,arr2:%v", arr1, arr2)
		merged := mergeOrderedArray3(arr1, arr2)
		t.Logf("after merge:%v", merged)
	})
}
