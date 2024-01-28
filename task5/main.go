package main

import (
	"fmt"
	"time"
)

func main() {
	// переменная start хранит время начала программы
	start := time.Now()
	n := 1
	ch := make(chan int)

	// горутина извлекает числа из канала ch
	go func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()

	// цикл выполняется пока с начала выполнения программмы прошло не более n секунд
	for i := 0; int(time.Since(start).Seconds()) < n; i++ {
		ch <- i
	}
	close(ch)
}
