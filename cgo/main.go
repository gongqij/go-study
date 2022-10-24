package main

import "C"

import (
	"fmt"

	_ "go-study/cgo/number"
)

func main() {
	println("Done")
}

//export goPrintln
func goPrintln(s *C.char) {
	fmt.Println("goPrintln:", C.GoString(s))
}
