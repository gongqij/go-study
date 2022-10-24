package main

import (
	"fmt"
	"go-study/go/happen-before/init/p1"
	_ "go-study/go/init/p4"
)

var (
	a = c + b // == 9
	b = f()   // == 4
	c = f()   // == 5
	d = 3     // == 5 全部初始化完成后
)

func f() int {
	d++
	return d
}

func init() {
	fmt.Println("init func in main")
}

func main() {
	fmt.Println(a, b, c, d)
	fmt.Println("P1_file0_V1:", p1.P1_file0_V1)
	fmt.Println("P1_file0_V2:", p1.P1_file0_V2)
}
