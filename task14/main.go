package main

import "fmt"

// использхуется оператор switch
func getType(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	}
	// с помощью функции Sprintf и спецификатора %T получаем тип интерфейса
	return fmt.Sprintf("%T\n", i)
}

func main() {
	fmt.Println(getType(123))
	fmt.Println(getType("123"))
	fmt.Println(getType(true))
	var c chan int
	fmt.Println(getType(c))
}
