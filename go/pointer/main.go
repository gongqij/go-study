package main

import "fmt"

type Student struct {
	name string
}

func main() {
	stu1 := &Student{name: "gongqi"}
	foo(stu1) //值传递
	fmt.Println(stu1)
	fmt.Printf("%p,%p\n", stu1, &stu1)
}

func foo(stu2 *Student) { //新的*Student类型的变量,值为地址
	fmt.Printf("%p,%p\n", stu2, &stu2)
	stu2 = nil //stu的值为地址，将地址置为nil
	fmt.Printf("%p,%p\n", stu2, &stu2)
}
