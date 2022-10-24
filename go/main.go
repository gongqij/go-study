package main

import "fmt"

func main() {
	var nilSlice []int
	emptySlice := make([]int, 5)
	fmt.Println(nilSlice)                     // Output: []
	fmt.Println(len(nilSlice), cap(nilSlice)) // Output: 0 0
	fmt.Println(nilSlice == nil)              // Output: true
	emptySlice = append(emptySlice, 1)
	fmt.Println(emptySlice)                       // Output: []
	fmt.Println(len(emptySlice), cap(emptySlice)) // Output: 0 0
	fmt.Println(emptySlice == nil)                // Output: false

}
