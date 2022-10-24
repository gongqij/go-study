package main

import "fmt"

func main() {
	//memoryAlignment()
	demo1()

}

func new() {
	s := []int{5, 6}
	s = append(s, 7, 8, 9)
	x := append(s, 11)
	y := append(s, 12)
	fmt.Printf("%p,%p,%p\n", &s, &x, &y) //切片本身的地址都不一样
	fmt.Printf("%p,%p,%p\n", s, x, y)    //切片指向数组地址一样，三个切片指向同一个底层数组
	fmt.Println(s, x, y)
}

func memoryAlignment() {
	s := []int{1, 2}
	s = append(s, 4, 5, 6)
	fmt.Printf("len=%d, cap=%d", len(s), cap(s))
}

func demo1() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	fmt.Printf("s1:len=%d,cap=%d\n", len(s1), cap(s1))
	//索引2（闭区间）到索引5（开区间，元素真正取到索引4），长度为3，容量默认到数组结尾，为8
	s2 := s1[2:6:7]
	fmt.Printf("s2:len=%d,cap=%d\n", len(s2), cap(s2))
	//s2 从 s1 的索引2（闭区间）到索引6（开区间，元素真正取到索引5），容量到索引7（开区间，真正到索引6），为5。
	s2 = append(s2, 100)
	s2 = append(s2, 200)
	s1[2] = 20
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(slice)
}

func demo2() {
	s := []int{1, 1, 1}
	newS := myAppend(s)

	fmt.Println(s)    //没改变s（首地址，长度，容量都没有变） 对s没有任何影响（底层数组也不会改变）
	fmt.Println(newS) //新切片

	s = newS

	myAppendPtr(&s)
	fmt.Println(s)
}

func myAppend(s []int) []int {
	// 这里 s 虽然改变了，但并不会影响外层函数的 s
	s = append(s, 100)
	return s
}

func myAppendPtr(s *[]int) {
	// 会改变外层 s 本身
	*s = append(*s, 100)
	return
}

func demo3() {
	s := []int{1, 1, 1}
	f(s)
	fmt.Println(s)
}

func f(s []int) {
	// i只是一个副本，不能改变s中元素的值
	/*for _, i := range s {
	      i++
	  }
	*/

	for i := range s {
		s[i] += 1
	}
}
