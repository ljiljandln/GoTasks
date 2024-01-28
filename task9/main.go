package main

import "fmt"

// функция пишет числа в первый канал конвейера
func source(numbers []int) <-chan int {
	// создается буферизованный канал с размером буфера равным длине numbers
	out := make(chan int, len(numbers))

	go func() {
		for _, num := range numbers {
			out <- num
		}
		// после обработки всех значений массива, канал закрывается
		close(out)
	}()
	// функция возвращает однонаправленный канал, доступный только для приема
	return out
}

// функция пишет во второй канал конвейера квадраты значений, полученных из первого канала
func square(in <-chan int) <-chan int {
	out := make(chan int, len(in))

	go func() {
		for num := range in {
			out <- num * num
		}
		close(out)
	}()
	return out
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// происходит итерация по возвращаемому значению функции square, которая в качестве агрумента
	// получает значения, возвращаемые функцией source
	for num := range square(source(numbers)) {
		fmt.Println(num)
	}
}
