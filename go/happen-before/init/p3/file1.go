package p3

import (
	"fmt"
	"go-study/go/happen-before/init/trace"
)

var P3_File1_v1 = trace.Trace("init P3_File1_v1", 3)
var P3_File1_v2 = trace.Trace("init P3_File1_v2", 3)

func init() {
	fmt.Println("init func in P3_File1")
	P3_File1_v1 = 300
	P3_File1_v2 = 300
}
