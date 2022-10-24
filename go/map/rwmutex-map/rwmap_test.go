package main

import (
	"fmt"
	"testing"
)

func TestRwMap(t *testing.T) {
	rwmap := NewRWMap(10)
	rwmap.Set("xxx", "xxxx")
}

func print(str string) {
	fmt.Println(str)

}
