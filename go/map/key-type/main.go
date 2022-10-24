package main

import "fmt"

type mapKey struct {
	key int
}

func main() {
	m1 := make(map[chan int]string)
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch1 <- 1
	ch2 <- 1
	fmt.Println(ch1 == ch2)
	//channel可以作为key
	m1[ch1] = "ch1"
	m1[ch2] = "ch2"

	var m = make(map[mapKey]string)
	var key = mapKey{10}

	m[key] = "hello"
	fmt.Printf("m[key]=%s\n", m[key])

	// 修改key的字段的值后再次查询map，无法获取刚才add进去的值
	key.key = 100
	fmt.Printf("再次查询m[key]=%s\n", m[key])
}
