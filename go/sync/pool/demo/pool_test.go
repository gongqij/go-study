package demo

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	sli := make([]int, 0, 10)
	sli2 := sli[:5]
	fmt.Println(sli2)
	for i := 1; i <= 100; i++ {
		pool(i)
	}
}

func BenchmarkPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pool(i)
	}
}

type Stu struct {
	name string
}

func Benchmark(b *testing.B) {
	b.ReportAllocs()
	var a, z [1000]*Stu
	p := sync.Pool{New: func() interface{} { return &Stu{} }}
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(a); j++ {
			a[j] = p.Get().(*Stu)
		}
		for j := 0; j < len(a); j++ {
			p.Put(a[j])
		}
		a = z
		runtime.GC()
	}
}
