package main

import (
	"fmt"
	"time"
)

func main() {
	echo(0)
	echo(int32(0))
	echo(uint32(0))
	echo(uint64(0))
	echo("hello")
	echo(struct{}{})
	echo(time.Now())
}

func echo[T any](t T) string {
	return fmt.Sprintf("%v", t)
}
