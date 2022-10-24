package main

import "fmt"

type T interface{}

var a T
var b *T
var c interface{} = a
var d interface{} = b

func main() {
	fmt.Println(a, b, c, d)
	fmt.Println(a == nil)         //a为T类型，未初始化，所以为nil
	fmt.Println(b == nil)         //b为指针类型*T，未初始化，所以为nil
	fmt.Println(c == a, c == nil) //c为interface类型，
	fmt.Println(d == b, d == nil) //d本身为interface类型，它存储的类型为*T，不满足类型与值同时为nil，所以d!=nil；d接口可以转换成b，所以可以比较，而且d接口的动态类型为*T，动态值为nil，所以与b相等
}

/*
nil,nil,nil,nil
true
true
true true
true false
*/
