package main

import "fmt"

func rangeMap() {
	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}

	counter := 0
	for k, v := range m {
		if counter == 0 {
			delete(m, "A")
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println(counter)
}

func appendSliceWhenRange() {
	arr := []int{1, 2, 3}
	for _, v := range arr { //range前预先获取了切片的长度
		arr = append(arr, v)
	}
	fmt.Println(arr)
}

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("original a =", a)

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}

	fmt.Println("after for range loop, r =", r)
	fmt.Println("after for range loop, a =", a)
}
