package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	n := 2000

	var count int32

	var ws sync.WaitGroup
	for i := 0; i < n; i++ {
		ws.Add(1)
		go func() {
			defer ws.Done()
			// используем атомарную операцию, чтобы избежать состояния гонки
			atomic.AddInt32(&count, 1)
		}()
	}

	ws.Wait()

	fmt.Println(count)
}
