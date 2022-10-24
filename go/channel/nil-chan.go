package main

import (
	"fmt"
	"time"
)

func sendToNilChan() {
	var ch chan int
	go func() {
		fmt.Println("向ni chan发送数据")
		ch <- 1
		fmt.Println("退出。。")
	}()
	time.Sleep(time.Second * 2)
	close(ch)
}

func recvFromClosedChan() {
	ch := make(chan int, 10)
	close(ch)
	res := <-ch
	fmt.Println(res)
}
