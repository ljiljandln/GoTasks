package main

import "fmt"

// возвращает индекс искомого элемента массива, если он есть, иначе возвращает -1
func binSearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		curr := arr[mid]
		if curr < target {
			left = mid + 1
		} else if curr > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 4, 56, 67}
	fmt.Printf("%d, %d", binSearch(arr, 56), binSearch(arr, 57))
}
