package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// Способ 1 - закрыть канал с помощью функции close
	ch := make(chan string)

	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			_, ok := <-ch
			if !ok {
				fmt.Println("G1 done")
				return
			}
		}
	}()

	ch <- "hello"
	ch <- "chan"
	close(ch)
	wg.Wait()

	// Способ 2 - создать канал done, в который для остановки будет передаваться пустая структура
	done := make(chan struct{})

	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			select {
			case <-done:
				fmt.Println("G2 done")
				return
			default:
				fmt.Println("G2 working")
			}
		}
	}()

	done <- struct{}{}
	wg.Wait()
}
