package removeDuplicates

import (
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	res := removeDuplicates([]int{2, 4, 5, 5, 7, 7})
	t.Logf("removeDuplicates:%v", res)
}
