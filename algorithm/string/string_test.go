package string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	assert.Equal(t, 1234, myAtoi("1234"))
	assert.Equal(t, -1234, myAtoi("-1234"))
	assert.Equal(t, +1234, myAtoi("+1234"))
}
