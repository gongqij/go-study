package main

import (
	"fmt"
	"runtime"
)

type User struct {
	ID     int64
	Name   string
	Avatar string
}

func GetUserInfo() int64 {
	u := &User{ID: 13746731, Name: "EDDYCJY", Avatar: "https://avatars0.githubusercontent.com/u/13746731"}
	return u.ID
}

func main() {
	//逃逸分析  go run -gcflags "-m -l" .
	str := new(string)
	*str = "EDDYCJY"
	//str作用域没有改变，未逃逸
	//fmt.Println(str) 加上这一行
	//会发生逃逸，str变量被分配到堆上
	//当形参为 interface 类型时，在编译阶段编译器无法确定其具体的类型。因此会产生逃逸，最终分配到堆上
	//如果你有兴趣追源码的话，可以看下内部的 reflect.TypeOf(arg).Kind() 语句，其会造成堆逃逸，而表象就是 interface 类型会导致该对象分配到堆上
	GetUserInfo()
	a := 2
	u := User{
		ID: int64(a),
	}
	fmt.Println(u)
	fmt.Println(runtime.NumCPU())
}
