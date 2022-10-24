package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type result struct {
	workerID    int
	sleepSecond int
}

const (
	workMaxDuration = 10
	timeout         = 3
)

func init() {
	rand.Seed(time.Now().Unix())
}

func (r result) String() string {
	return fmt.Sprintf("worker %d done in %ds", r.workerID, r.sleepSecond)
}

func main() {
	workerNum := 5
	data := make(chan result, workerNum)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*timeout)
	defer cancel()
	for i := 0; i < workerNum; i++ {
		go func(ctx context.Context, index int) {
			resChan := make(chan result, 1)
			go func(index int) {
				res := Serve(index)
				resChan <- res
			}(index)
			select {
			case <-ctx.Done():
				fmt.Printf("第%d个worker未执行完成，超时退出\n", index)
				return
			case res := <-resChan:
				data <- res
				fmt.Printf("第%d个worker执行完成，退出\n", index)
			}
		}(ctx, i)
	}
	for {
		select {
		case <-time.After(time.Second * workMaxDuration):
			return
		case res := <-data:
			fmt.Println(res)
		}
	}
}

func Serve(index int) result {
	n := rand.Intn(workMaxDuration)
	fmt.Printf("第%d个worker执行完成需要%ds...\n", index, n)
	time.Sleep(time.Second * time.Duration(n))
	return result{
		workerID:    index,
		sleepSecond: n,
	}
}
