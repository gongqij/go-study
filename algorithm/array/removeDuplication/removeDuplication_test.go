package removeDuplication

import (
	"fmt"
	"testing"
)

func TestOrderedArrayRemoveDuplicates(t *testing.T) {
	nums := []int{0, 1, 1, 2, 2, 2, 3}
	fmt.Println(orderedArrayRemoveDuplicates(nums))
}

func TestUnorderedArrayRemoveDuplication(t *testing.T) {
	nums := []int{4, 1, 2, 2, 3, 2, 5}
	fmt.Println(unorderedArrayRemoveDuplication(nums))
}
