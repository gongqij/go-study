package main

import (
	"fmt"
	"sync"
)

func demo1() {
	var once sync.Once
	once.Do(func() {
		once.Do(func() {
			fmt.Println("初始化")
		})
	})
}

type MuOnce struct {
	sync.Once
}

func (m *MuOnce) Reset() {
	m.Once = sync.Once{}
}

func demo2() {
	var m MuOnce
	m.Do(m.Reset)
}
