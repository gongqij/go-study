package _defer

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	fmt.Println("test1...")               // 打印结果为 return: 2
	fmt.Println("test1 return:", test1()) // 打印结果为 return: 2
	fmt.Println("test2...")               // 打印结果为 return: 2
	fmt.Println("test2 return:", test2()) // 打印结果为 return: 2
	fmt.Println("test3...")               // 打印结果为 return: 2
	a, b := test3()
	fmt.Println("test3 return:", a, *b) // 打印结果为 return: 2
}

func TestTest4(t *testing.T) {
	test4()
}
