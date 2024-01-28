package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	arr := []int32{2, 4, 6, 8, 10}

	// для того, чтобы подождать выполнение всеъ горутин используем sync.WaitGroup
	var wg sync.WaitGroup
	var res int32

	for _, val := range arr {
		// добавляем горутину
		wg.Add(1)
		add := val * val
		go func() {
			// по завершении работы функции закрываем горутину
			defer wg.Done()
			// используем атомарную операцию
			atomic.AddInt32(&res, add)
		}()
	}

	// ждем завершение всех горутин
	wg.Wait()

	fmt.Println(res)
}
