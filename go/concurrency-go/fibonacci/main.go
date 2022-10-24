package main

import "fmt"

func main() {
	fmt.Printf("fib(4) = %d\n", <-fibonacci(4))
	for i := 0; i < 10; i++ {
		fmt.Print(<-fibonacci(i), " ")
	}
	fmt.Println()
	for i := 0; i < 10; i++ {
		fmt.Print(fibonacci2(i), " ")
	}
	fmt.Println()
	f := fibonacci3()
	for i := 0; i < 10; i++ {
		fmt.Print(f(), " ")
	}
}

func fibonacci(n int) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		if n < 2 {
			result <- 1
			return
		}
		result <- <-fibonacci(n-1) + <-fibonacci(n-2)
	}()
	return result
}

func fibonacci2(num int) int {
	if num < 2 {
		return 1
	}
	return fibonacci2(num-1) + fibonacci2(num-2)
}

func fibonacci3() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func fibonacci4(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
