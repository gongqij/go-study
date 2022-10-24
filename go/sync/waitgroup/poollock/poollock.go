package poollock

import (
	"fmt"
	"sync"
)

// PoolLock ...
type PoolLock struct {
	wg *sync.WaitGroup
	ch chan byte
}

// Init ...
func (p *PoolLock) Init(maxCount int) {
	p.wg = new(sync.WaitGroup)
	p.ch = make(chan byte, maxCount)
}

// Add ...
func (p *PoolLock) Add() {
	if p.wg == nil {
		return
	}
	p.wg.Add(1)
	p.ch <- 0
}

// Wait ...
func (p *PoolLock) Wait() {
	if p.wg == nil {
		fmt.Println("PoolLock not initialized")
		return
	}
	p.wg.Wait()
}

// Done ...
func (p *PoolLock) Done() {
	if p.wg == nil {
		fmt.Println("PoolLock not initialized")
		return
	}
	<-p.ch
	p.wg.Done()
}
