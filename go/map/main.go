package main

import "fmt"

type coder interface {
	code()
	debug()
}

type Gopher struct {
	Language string
}

func (p Gopher) code() { // 1、隐式实现 func (p *Gopher) code()
	fmt.Printf("I am coding %s language\n", p.Language)
}

func (p *Gopher) debug() { // 2、没有实现 func (p Gopher) debug()
	fmt.Printf("I am debuging %s language\n", p.Language)
}

func main() {
	var c coder = &Gopher{"Go"} //不报错，因为原因1；去掉&报错，因为原因2
	c.code()
	c.debug()
}
