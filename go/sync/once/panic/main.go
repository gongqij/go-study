package main

import (
	"fmt"
	"sync"
)

func main() {
	var on sync.Once
	defer func() {
		recover()
	}()
	on.Do(test)
	on.Do(test)
}

func test() {
	fmt.Println("before")
	panic("xxx")
}
