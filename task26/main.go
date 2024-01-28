package main

import (
	"fmt"
	"strings"
)

// для проверки на уникальность тспользуется отображение
func isUniq(str string) bool {
	str = strings.ToUpper(str)
	charsMap := make(map[rune]bool)

	// итерация по символам строки
	for _, ch := range str {
		// если по ключу получаем значение true, значит символ уже встречался в строке
		// таким образом возвращаем false в качестве результата
		if charsMap[ch] {
			return false
		} else {
			// иначе кладем значение true для ключа ch в отображение
			charsMap[ch] = true
		}
	}

	// если дошли до этой строки, значит проитерировались по всем символам строки
	// не выходя из цикла, возвращаем true
	return true
}

func main() {
	fmt.Printf("%t, %t", isUniq("abCd"), isUniq("Aacd"))
}
