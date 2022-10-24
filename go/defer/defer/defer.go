package _defer

import (
	"errors"
	"fmt"
	"time"
)

func test1() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer: 2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer: 1
	}()
	return i
}

func test2() (err error) {
	defer func() {
		err = errors.New("A")
		fmt.Println(err) //A  第二个defer输出A，最后返回值被赋值为A
	}()
	defer func(err error) {
		err = errors.New("B")
		fmt.Println(err) // 第一个defer输出B
	}(err) //不会修改返回值err
	return errors.New("C") //返回值C被覆盖了
}
func test3() (int, *int) {
	var i int
	defer func() {
		i++
		fmt.Println("c defer2:", i) // 第二个defer输出,打印结果为 c defer: 2
	}()
	defer func() {
		i++
		fmt.Println("c defer1:", i) // 第一个defer输出,打印结果为 c defer: 1

	}()
	return i + 1, &i
}

//调用 defer 关键字会立刻拷贝函数中引用的外部参数，所以 time.Since(startedAt) 的结果不是在 main 函数退出之前计算的，
//而是在 defer 关键字调用时计算的，最终导致上述代码输出 0s。
func test4() {
	startedAt := time.Now()
	defer fmt.Println(time.Since(startedAt))

	time.Sleep(time.Second)
}
