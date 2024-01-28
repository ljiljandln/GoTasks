package main

import (
	"fmt"
	"strings"
)

func reverse(str string) string {
	/* получаем срез строк с помощью функции пакета strings, которая делит
	строку по пробельным символам */
	words := strings.Fields(str)

	// используем два указателя в цикле
	left, right := 0, len(words)-1
	for left < right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}

	// соединяем перевернутые слова с помощью функции Join
	s := strings.Join(words, " ")
	return s
}

func main() {
	str := "one two three four"
	str = reverse(str)
	fmt.Print(str)
}
