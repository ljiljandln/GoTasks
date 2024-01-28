package main

import "fmt"

func reverse(str string) string {
	// преобразуем str в срез рун для того, чтобы символы состоящие из более чем
	//одного бита, обрабатывались корректно
	arr := []rune(str)
	// используем два указателя, пока левый ментше правого, меняя значения по этим индексам
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	// выполняем обратное преобразование из среза рун в строку, и возвращаем результат
	return string(arr)
}

func main() {
	str := "абвгд"

	fmt.Print(reverse(str))
}
