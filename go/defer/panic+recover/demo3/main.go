package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("first defer")
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
			defer func() {
				if err := recover(); err != nil {
					fmt.Println(err)
				}
			}()
			panic("panic again and again")
		}()
		panic("panic again")
	}()

	panic("panic")
}
