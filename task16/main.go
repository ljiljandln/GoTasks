package main

import "fmt"

/*
Функция делит массив на три части: меньше чем pivot, равную ему и больше него.
Возвращает индекс первого элемента, равного pivot, и индекс первого элемента,
большего чем pivot.
Таким образом по завершению работы функции нужно отсортировать как минимум на один
элемент меньше, чем вначале
*/
func partition(left, right int, arr []int) (equal, greater int) {
	index := (left + right) / 2
	pivot := arr[index]

	equal, greater = left, left

	for now := left; now <= right; now++ {
		if arr[now] < pivot {
			tmp := arr[now]
			arr[now] = arr[greater]
			arr[greater] = arr[equal]
			arr[equal] = tmp
			greater++
			equal++
		} else if arr[now] == pivot {
			arr[greater], arr[now] = arr[now], arr[greater]
			greater++
		}
	}
	return
}

/*
Функция вызывается, когда длина неотсортированного участка массива меньше 16.
В этом случае сортировка вставками работает быстрее, чем рекурсивная qSort с вызовом
partition, особенно на частично отсортированном массиве
*/
func insertionSort(left, right int, arr []int) {
	for i := left + 1; i <= right; i++ {
		k := i
		tmp := arr[k]
		for k > left && tmp < arr[k-1] {
			arr[k] = arr[k-1]
			k--
		}
		arr[k] = tmp
	}
}

func quickSort(left, right int, arr []int) {
	for left < right {
		if right-left > 16 {
			equal, greater := partition(left, right, arr)
			/* После отработки функции partition, qSort рекурсивно вызывается для меньшей части
			неотсортированного массива. Большая часть продолжает сортироваться в цикле.
			Таким образом уменьшается количество рекурсий.*/
			if equal-left-1 < right-greater {
				quickSort(left, equal-1, arr)
				left = greater
			} else {
				quickSort(greater, right, arr)
				right = equal - 1
			}
		} else {
			insertionSort(left, right, arr)
			return
		}
	}
}

func main() {
	arr := make([]int, 100)

	curr := 100
	for i := range arr {
		arr[i] = curr
		if i%2 == 1 {
			curr--
		}
	}

	quickSort(0, 99, arr)

	fmt.Print(arr)
}
