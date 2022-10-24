package p3

import (
	"fmt"
	"go-study/go/happen-before/init/trace"
)

var P3_File0_V1 = trace.Trace("init P3_File0_V1", 3)

func init() {
	fmt.Println("init func in P3_File0")
}
