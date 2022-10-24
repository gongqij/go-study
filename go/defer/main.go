package main

import "fmt"

func main() {
	defer func() {
		defer func() {
			if err := recover(); err != nil { //无法recover‘panic again’，因为隔了两层defer
				fmt.Println(err)
			}
		}()
	}()
	defer func() {
		if err := recover(); err != nil { //可以recover‘panic once’
			fmt.Println(err)
		}
		defer func() {
			defer func() {
				if err := recover(); err != nil { ////可以recover‘panic again and again’
					fmt.Println(err)
				}
			}()
			panic("panic again and again")
		}()
		panic("panic again") //无法被recover
	}()
	panic("panic once")
}
