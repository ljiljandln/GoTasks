package main

import "fmt"

func main() {
	a, b := 5, 6

	// используется множественное присваивание
	a, b = b, a

	fmt.Printf("a = %d, b = %d", a, b)
}
