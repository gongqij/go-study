package qsort

import (
	"fmt"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {

	values := []int32{42, 9, 101, 95, 27, 25}
	Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	fmt.Println("sort by C.qsort:", values)

	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	fmt.Println("sort by go.sort.Slice:", values)
}
