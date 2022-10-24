package typeConvert

/*
 #include <string.h>
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

func CArrayToGoSlice(cArray *C.char) {
	// 通过 reflect.SliceHeader 转换
	var arr0 []byte
	var arr0Hdr = (*reflect.SliceHeader)(unsafe.Pointer(&arr0))
	arr0Hdr.Data = uintptr(unsafe.Pointer(cArray))
	arr0Hdr.Len = 10
	arr0Hdr.Cap = 10
	fmt.Println("通过reflect.SliceHeader转换,arr0=", arr0)

	// 通过切片语法转换
	arr1 := (*[31]byte)(unsafe.Pointer(cArray))[:10:10]
	fmt.Println("通过切片语法转换,arr1=", arr1)
}

func CStringToGoString(cString *C.char) {
	var s0 string
	var s0Hdr = (*reflect.StringHeader)(unsafe.Pointer(&s0))
	s0Hdr.Data = uintptr(unsafe.Pointer(cString))
	s0Hdr.Len = int(C.strlen(cString))
	fmt.Println("通过reflect.StringHeader转换,s0:", s0)

	sLen := int(C.strlen(cString))
	s1 := string((*[31]byte)(unsafe.Pointer(cString))[:sLen:sLen])
	fmt.Println("通过切片语法转换,s1:", s1)
}

func PointToPointConvert() {
	//unsafe.Pointer指针类型类似C语言中的void * 类型的指针
	//任何类型的指针都可以通过强制转换为unsafe.Pointer指针类型去掉原有的类型信息，然后再重新赋予新的指针类型而达到指针间的转换的目的。
	var p *int
	var q *float64

	q = (*float64)(unsafe.Pointer(p)) // *int => *float64
	p = (*int)(unsafe.Pointer(q))     // *float64 => *int
}

func ValueToPointConvert() {
	p := 1

	intPoint := (*C.char)(unsafe.Pointer(uintptr(p))) // int => *int
	fmt.Println(intPoint)
}

func SliceToSlice() {
	var p []int64
	var q []int32

	pHdr := (*reflect.SliceHeader)(unsafe.Pointer(&p))
	qHdr := (*reflect.SliceHeader)(unsafe.Pointer(&q))

	pHdr.Data = qHdr.Data
	//pHdr.Len = qHdr.Len * unsafe.Sizeof(q[0]) / unsafe.Sizeof(p[0])
	//pHdr.Cap = qHdr.Cap * unsafe.Sizeof(q[0]) / unsafe.Sizeof(p[0])
}
