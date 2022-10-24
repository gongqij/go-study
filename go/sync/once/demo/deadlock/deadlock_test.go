package main

import "testing"

func TestDeadlock(t *testing.T) {
	demo1()
	demo2()
}
