package main

import "fmt"

//必须在defer函数中直接调用recover，不能进行封装或者嵌套
func wrongDemo() {
	defer func() {
		if r := MyRecover(); r != nil {
			fmt.Println(r)
		}
	}()
	panic("err")
}
func MyRecover() interface{} {
	fmt.Println("recover")
	return recover()
}

//必须要和有异常的栈帧只隔一个栈帧， recover 函数才能正常捕获异常。
//换言之， recover 函数捕获的是祖父一级调用函数栈帧的异常（刚好可以跨越一层 defer 函数）！
func wrongDemo2() {
	defer func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()
	}()
	panic("err")
}
