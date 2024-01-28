package main

import (
	"fmt"
	"sync"
)

// ConcurrentMap отображение с полем mutex
type ConcurrentMap struct {
	mp    map[uint8]uint8
	mutex sync.Mutex
}

func (cp *ConcurrentMap) add(num uint8) {
	// блокируем mutex, доступ будет только у одной горутины
	cp.mutex.Lock()
	// после завершения работы функции снимаем блокировку
	defer cp.mutex.Unlock()
	cp.mp[num]++
}

func getSlice() []uint8 {
	nums := make([]uint8, 100)
	for i, k := 0, uint8(0); i < len(nums); i++ {
		if i%4 == 0 {
			k++
		}
		nums[i] = k
	}
	return nums
}

func main() {
	workers := 4
	nums := getSlice()

	jobs := make(chan uint8, 10)
	go func() {
		for _, num := range nums {
			jobs <- num
		}
		close(jobs)
	}()

	cm := ConcurrentMap{mp: make(map[uint8]uint8)}

	// wg используется для ожидания завершения группы Goroutine
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for num := range jobs {
				cm.add(num)
			}
		}()
	}

	wg.Wait()

	fmt.Println(cm.mp)
}
