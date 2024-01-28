package main

import (
	"fmt"
	"math/big"
)

func add(x, y *big.Int) *big.Int {
	return new(big.Int).Add(x, y)
}

func sub(x, y *big.Int) *big.Int {
	return new(big.Int).Sub(x, y)
}

func mul(x, y *big.Int) *big.Int {
	return new(big.Int).Mul(x, y)
}

func div(x, y *big.Int) *big.Int {
	return new(big.Int).Div(x, y)
}

func main() {
	// с помощью пакета big, создаем два больших целых числа
	x, _ := new(big.Int).SetString("10000000000000000000000", 10)
	y, _ := new(big.Int).SetString("500000000000000000000", 10)

	fmt.Println(add(x, y))
	fmt.Println(sub(x, y))
	fmt.Println(mul(x, y))
	fmt.Println(div(x, y))
}
