package main

import "fmt"

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}

	// создаем множество
	set := make(map[string]bool, 5)
	// итерируемся по значениям в words, добавляем каждое в set
	for _, word := range words {
		set[word] = true
	}

	fmt.Print(set)
}
