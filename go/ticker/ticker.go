package main

import (
	"fmt"
	"time"
)

func main() {
	tic := time.NewTicker(time.Second * 2)
	defer tic.Stop()
	go func() {
		for {
			select {
			case <-tic.C:
				fmt.Println(1, time.Now())
			}
		}
	}()
	go func() {
		for {
			select {
			case <-tic.C:
				fmt.Println(2, time.Now())
			}
		}
	}()
	select {}
}
