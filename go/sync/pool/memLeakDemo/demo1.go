package main

import (
	"bytes"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main1() {
	pool := sync.Pool{New: func() interface{} {
		return new(bytes.Buffer)
	}}

	processRequest := func(size int) {
		b := pool.Get().(*bytes.Buffer)
		time.Sleep(500 * time.Millisecond) // Simulate processing time
		b.Grow(size)
		pool.Put(b)
		time.Sleep(1 * time.Millisecond) // Simulate idle time
	}

	// Simulate a set of initial large writes.
	for i := 0; i < 10; i++ {
		go func() {
			processRequest(1 << 28) // 256MiB
		}()
	}

	time.Sleep(time.Second) // Let the initial set finish

	// Simulate an un-ending series of small writes.
	for i := 0; i < 10; i++ {
		go func() {
			for {
				processRequest(1 << 10) // 1KiB
			}
		}()
	}

	// Continually run a GC and track the allocated bytes.
	var stats runtime.MemStats
	for i := 0; ; i++ {
		runtime.ReadMemStats(&stats)
		fmt.Printf("Cycle %d: %dB\n", i, stats.Alloc)
		time.Sleep(time.Second)
		runtime.GC()
	}
}
