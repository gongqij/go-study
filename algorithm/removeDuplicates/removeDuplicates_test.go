package removeDuplicates

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"sort"
)

func Test_removeDuplicates(t *testing.T) {
	res1 := ascArrayRemoveDuplication([]int{2, 4, 5, 5, 7, 7})
	assert.Equal(t, res1, []int{2, 4, 5, 7})
	res2 := arrayRemoveDuplication1([]int{2, 4, 5, 5, 7, 7})
	sort.Ints(res2)
	assert.Equal(t, res2, []int{2, 4, 5, 7})
	res3 := arrayRemoveDuplication2([]int{5, 2, 5, 2, 3, 4, 5, 1})
	sort.Ints(res3)
	assert.Equal(t, res3, []int{1, 2, 3, 4, 5})
}
