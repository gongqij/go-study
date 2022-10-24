package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type workerChan struct {
	lastUseTime time.Time
	ch          chan int
}

var workerBuffers = sync.Pool{
	New: func() interface{} {
		return &workerChan{
			lastUseTime: time.Now(),
			ch:          make(chan int, 100),
		}
	},
}

func main() {
	vch1 := workerBuffers.Get()
	ch1 := vch1.(*workerChan)
	ch1.ch <- 999
	fmt.Printf("%p,lastUseTime:%v\n", ch1, ch1.lastUseTime)
	workerBuffers.Put(vch1)
	runtime.GC()
	vch2 := workerBuffers.Get()
	ch2 := vch2.(*workerChan)
	fmt.Printf("%p,lastUseTime:%v,len:%d\n", ch2, ch2.lastUseTime, len(ch2.ch))
	if ch1.lastUseTime == ch2.lastUseTime {
		fmt.Println("两次使用了同一个")
	}
}
