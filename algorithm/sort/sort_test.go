package sort

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"log"
)

var input = []int{5, 2, 3, 1, 4}
var output = []int{1, 2, 3, 4, 5}

func TestBubbleSort(t *testing.T) {
	log.Printf("before sort:%v", input)
	res := BubbleSort(input)
	log.Printf("after  sort:%v", res)
	assert.Equal(t, true, reflect.DeepEqual(output, res))
}
