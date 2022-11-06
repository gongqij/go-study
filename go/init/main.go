package main

import (
	"fmt"
	"select/a"
)

func main() {
	fmt.Println("main running")
	a.Print()
}

func init() {
	fmt.Println("xxx")
}
