package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []int{5, 2, 3, 1, 4}
var output = []int{1, 2, 3, 4, 5}

func TestBubbleSort(t *testing.T) {
	assert.Equal(t, output, BubbleSort(input), "BubbleSort")
	assert.Equal(t, output, MergeSort(input), "MergeSort")
}
