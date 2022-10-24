package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func runPlay() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer func() {
			fmt.Println("---------------------")
			fmt.Println("号令枪打响，开始比赛。。。")
			fmt.Println("---------------------")
			wg.Done()
		}()
		time.Sleep(time.Second)
	}()
	ch := make(chan int)
	playerNum := 5
	for i := 1; i <= playerNum; i++ {
		go func(i int) {
			fmt.Println("选手", i, "号等待号令枪响。。。")
			wg.Wait()
			start := time.Now()
			fmt.Println("选手", i, "号开始奔跑。。。")
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
			dur := time.Since(start)
			fmt.Println("选手", i, "号到达终点。。。用时", dur.Seconds(), "s")
			ch <- i
		}(i)
	}
	rank := make([]int, 0, 5)
	for num := range ch {
		rank = append(rank, num)
		if len(rank) == playerNum {
			break
		}
	}
	fmt.Println("最终排名为:", rank)
}
