package main

import (
	"fmt"
	"testing"
	"time"
)

func TestChQueue(t *testing.T) {
	cq := NewChQueue(2)
	go func() {
		for {
			time.Sleep(time.Second * 2)
			cq.Enqueue(1)
		}
	}()
	go func() {
		for {
			fmt.Println("等待出队。。。")
			fmt.Println("出队", cq.Dequeue())
		}
	}()
	select {}
}

func TestCondQueue(t *testing.T) {
	q := NewQueue(2)
	go func() {
		for {
			time.Sleep(time.Second * 2)
			q.Enqueue(1)
		}
	}()
	go func() {
		for {
			fmt.Println("等待出队。。。")
			fmt.Println("出队", q.Dequeue())
		}
	}()
	select {}
}
