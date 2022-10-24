package demo2

import "fmt"

//发生panic后，依次从defer链表头获取deferFunc执行
func deferTest1() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer fmt.Println("first defer")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
