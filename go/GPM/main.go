package main

import (
	"net/http"
	"sync"

	"log"
	_ "net/http/pprof"
)

func main() {
	exitCh := make(chan int)
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			http.Get(`https://httpstat.us/200?sleep=10000`)

			wg.Done()
		}()
	}
	wg.Wait()
	<-exitCh
}
