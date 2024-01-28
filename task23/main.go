package main

import "fmt"

// удаляет элемент с сохранением исходного порядка
func deleteElement(slice []int, i int) []int {
	// при некорректных значениях i возвращаем исходный срез
	if i >= len(slice) || i < 0 {
		return slice
	}
	// соединяем два среза: до iого и от i+1 до конца
	return append(slice[:i], slice[i+1:]...)
}

// удаляет элемент без сохранения исходного порядка
func deleteElement1(slice []int, i int) []int {
	if i >= len(slice) || i < 0 {
		return slice
	}
	tmp := slice[len(slice)-1]
	slice[i] = tmp
	return slice[:len(slice)-1]
}

func main() {
	slice := []int{1, 2, 3, 3, 4, 4, 5}
	slice = deleteElement(slice, 2)
	fmt.Println(slice)
	slice = deleteElement1(slice, 3)
	fmt.Println(slice)
}
