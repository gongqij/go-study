package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		defer func() {
			fmt.Println("defer before f")
			if err := recover(); err != nil { //捕获到panic2
				//recover1
				fmt.Println("recover before f: ", err)
			}
			c <- 1
		}()
		f()
	}()
	<-c
	fmt.Println("main done")
}

func f() {
	fmt.Println("In f") //首先打印
	g()
	fmt.Println("Exit f")
}

func g() {
	defer func() {
		fmt.Println("defer in g")
		if err := recover(); err != nil { //捕获到panic1
			//recover2
			fmt.Println("recover in g: ", err)
		}
		//panic2
		panic("panic in g defer") //会被recover1捕获到
	}()
	//panic1
	panic("Panic in g")
	fmt.Println("Exit g") //不会被执行
}
