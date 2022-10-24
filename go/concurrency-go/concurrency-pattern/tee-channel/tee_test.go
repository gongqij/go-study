package main

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	done := make(chan interface{})
	defer close(done)
	for v := range repeat(done, 1, 2, 3) {
		fmt.Println(v)
	}
}

func TestTee(t *testing.T) {
	done := make(chan interface{})
	defer close(done)
	out1, out2 := tee(done, take(done, repeat(done, 1, 2), 4))
	for val1 := range out1 {
		fmt.Printf("out1: %v, out2: %v\n", val1, <-out2)
	}
}
