package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

// Фунцкия для запуска необходимого количества воркеров.
func startWorkers(n int, in <-chan int, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				num, ok := <-in
				if !ok {
					return
				}
				fmt.Println(num)
			}
		}()
	}
}

func getWorkersCount() int {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil || n <= 0 {
		fmt.Println("Неверный формат, введите положительное целое число")
		return getWorkersCount()
	}
	return n
}

func main() {
	// Cоздание канала для сигналов ОС.
	syscallChan := make(chan os.Signal, 1)
	signal.Notify(syscallChan, os.Interrupt)

	fmt.Println("Введите число воркеров")
	n := getWorkersCount()
	// Канал для постоянной записи данных.
	in := make(chan int, n*2)

	var wg sync.WaitGroup
	// Запуск воркеров
	startWorkers(n, in, &wg)

	for i := 0; ; i++ {
		select {
		case <-syscallChan:
			// канал закрывается, ждем когда все горутины закончат работу
			close(in)
			wg.Wait()
			return
		case in <- i:
		}
	}
}
