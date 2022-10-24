package p2

import (
	"fmt"
	"go-study/go/happen-before/init/p3"
	"go-study/go/happen-before/init/trace"
)

var P2_File0_v1 = trace.Trace("init P2_File0_v1", 2)
var P2_File0_v2 = trace.Trace("init P2_File0_v2", p3.P3_File1_v2)

func init() {
	fmt.Println("init func in P2_File0")
	P2_File0_v1 = 200
}
