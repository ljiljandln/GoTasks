package main

import "fmt"

type Human struct {
	Name string
	Age  uint8
}

// Action в структуру встроено поле типа Human
type Action struct {
	Human
	ActionType string
}

func main() {
	ac := Action{Human{"Holden", 16}, "traveling"}
	fmt.Println(ac.Name)
}
