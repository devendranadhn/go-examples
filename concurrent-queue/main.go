package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup
var wgD sync.WaitGroup

type ConcurrentQueue struct {
	queue []int32
	mu    sync.Mutex
}

func (q *ConcurrentQueue) Enqueue(item int32) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.queue = append(q.queue, item)
}

func (q *ConcurrentQueue) Dequeue() int32 {

	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.queue) == 0 {
		panic("Cannot deque a value from empty queue")
	}

	item := q.queue[0]
	q.queue = q.queue[1:]
	return item
}

func (q *ConcurrentQueue) Size() int {
	return len(q.queue)
}

func main() {
	q1 := ConcurrentQueue{
		queue: make([]int32, 0),
	}

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			q1.Enqueue(rand.Int31())
			wg.Done()
		}()
	}

	for i := 0; i < 100000; i++ {
		wgD.Add(1)
		go func() {
			q1.Dequeue()
			wgD.Done()
		}()
	}
	wg.Wait()
	wgD.Wait()
	fmt.Println(q1.Size())
}
