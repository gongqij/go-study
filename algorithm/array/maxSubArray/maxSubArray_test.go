package maxSubArray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	assert.Equal(t, maxSubArray([]int{-1, 2, -1, 3, -4}), 4)
}
