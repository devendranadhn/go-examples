package main

import (
	"fmt"
	"sync"
)

var count int = 0
var wg sync.WaitGroup
var mu sync.Mutex

func incCount() {
	mu.Lock()
	count++
	mu.Unlock()
	wg.Done()
}

func doCount() {
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go incCount()
	}
}

func main() {

	count = 0

	doCount()

	wg.Wait()

	fmt.Println("Final value of count is", count)

}
