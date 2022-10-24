package main

import (
	"fmt"
	"unsafe"
)

func IsLittleEndian() bool {
	var value int32 = 1 // 占4byte 转换成16进制 0x00 00 00 01
	// 大端(16进制)：00 00 00 01
	// 小端(16进制)：01 00 00 00
	pointer := unsafe.Pointer(&value)
	pb := (*[4]byte)(pointer)
	fmt.Println(&pb[0])
	fmt.Println(&pb[1])
	fmt.Println(&pb[2])
	fmt.Println(&pb[3])
	pb2 := (*byte)(pointer)
	fmt.Println(pb2)
	if *pb2 != 1 {
		return false
	}
	return true
}

func main() {
	fmt.Println(IsLittleEndian())
}

// 运行结果：ture
