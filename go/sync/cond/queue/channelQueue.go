package main

type ChQueue struct {
	data chan interface{}
}

func NewChQueue(capacity int) *ChQueue {
	return &ChQueue{data: make(chan interface{}, capacity)}
}

func (q *ChQueue) Enqueue(d interface{}) {
	q.data <- d
}

func (q *ChQueue) Dequeue() interface{} {
	d := <-q.data
	return d
}
