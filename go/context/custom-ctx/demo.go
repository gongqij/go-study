package main

import "fmt"

type customContext interface {
	Done()
}

type cancelCtx struct {
	customContext
	name string
}

func (s *cancelCtx) Done() {
	fmt.Println(s.name)
}

type valueCtx struct {
	customContext
	key, value string
}

func main() {
	canCtx1 := &cancelCtx{name: "wang"}
	canCtx2 := &cancelCtx{customContext: canCtx1, name: "li"}
	canCtx3 := &cancelCtx{customContext: canCtx2, name: "zhang"}
	valCtx1 := &valueCtx{canCtx3, "valCtx1", "yyy"}
	valCtx2 := &valueCtx{valCtx1, "valCtx2", "yyy"}
	valCtx2.Done() //TODO 为什么这里调的是canCtx3的Done
}
