package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("第一个goroutine开始接收")
		<-ch
		fmt.Println("第一个goroutine获取到")
	}()
	time.Sleep(time.Second)
	go func() {
		fmt.Println("第二个goroutine开始接收")
		<-ch
		fmt.Println("第二个goroutine获取到")
	}()
	time.Sleep(time.Second)
	ch <- 1
	time.Sleep(time.Second)
}
