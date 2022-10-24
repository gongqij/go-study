package main

import "fmt"

func rangeMap() {
	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}

	counter := 0
	for k, v := range m { //原map
		if counter == 0 {
			delete(m, "A")
		}
		counter++
		fmt.Println(k, v)
	}
	/*
		range第一次获取到A，那么打印为:3
		range第一次获取到B或C，那么打印为:2
	*/
	fmt.Println(counter)
}

func rangeSlice() {
	var a = []int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("original a =", a)

	/*
		a为数组，for range时copy一份数组aa，a[1]不会影响到aa，因为底层是两个不同的数组
		a为切片，for range时copy一份切片aa，a[1]可以影响到aa，因为底层共用一个数组
	*/
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

//对于所有的 range 循环，Go 语言都会在编译期将原切片或者数组赋值给一个新变量 ha，
//在赋值的过程中就发生了拷贝，而我们又通过 len 关键字预先获取了切片的长度，
//所以在循环中追加新的元素也不会改变循环执行的次数
func appendSliceWhenRange() {
	arr := []int{1, 2, 3}
	for _, v := range arr { //range前预先获取了切片的长度
		arr = append(arr, v)
	}
	fmt.Println(arr)
}

func getPointerWhenRange() {
	//错误的方式
	//输出333
	arr := []int{1, 2, 3}
	newArr := []*int{}
	for _, v := range arr {
		//Go语言会额外创建一个新的 v2 变量存储切片中的元素，循环中使用的这个变量v2会在每一次迭代被重新赋值而覆盖，赋值时也会触发拷贝。
		//在循环中获取返回变量v的地址都完全相同
		newArr = append(newArr, &v)
	}
	for _, v := range newArr {
		fmt.Println(*v)
	}

	//正确的方式
	//输出123
	arr2 := []int{1, 2, 3}
	newArr2 := []*int{}
	for i, _ := range arr2 {
		newArr2 = append(newArr2, &arr2[i])
	}
	for _, v := range newArr2 {
		fmt.Println(*v)
	}
}
