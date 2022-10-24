package p1

import (
	"fmt"
	"go-study/go/happen-before/init/p2"
	"go-study/go/happen-before/init/trace"
)

var P1_file0_V1 = trace.Trace("init P1_file0_V1", p2.P2_File0_v1)
var P1_file0_V2 = trace.Trace("init P1_file0_V2", p2.P2_File0_v2)

func init() {
	fmt.Println("init func in P1_file0")
}
