package demo

import (
	"testing"
)

type Student struct {
	name string
	age  int
}

func TestLockFreeQ(t *testing.T) {
	stu1 := Student{
		name: "gq",
		age:  29,
	}
	stu2 := Student{
		name: "gh",
		age:  31,
	}
	stu3 := Student{
		name: "gl",
		age:  34,
	}
	lkq := NewLKQueue()
	lkq.Enqueue(stu1)
	lkq.Enqueue(stu2)
	lkq.Enqueue(stu3)
	t.Log("入队:", stu1, stu2, stu3)
	t.Log("出队:", lkq.Dequeue(), lkq.Dequeue(), lkq.Dequeue())
}
