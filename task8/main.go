package main

import "fmt"

func setBit(number int64, i, bit uint) int64 {
	/*
		Если переданное значение i некорректно, возвращаем само число
	*/
	if i > 63 || i < 0 {
		return number
	}

	if bit == 0 {
		/* если нужно установить в iом бите значение 0, создаем битовую маску
		с 0 в iом бите и с 1 во всех остальных */
		var mask int64 = ^(1 << i)
		/* применяем операцию побитового И с приравниваем, все биты исходного числа,
		кроме iго не поменяют свое значение */
		number &= mask
	} else {
		/*
			используем побитовое ИЛИ с приравниванием
		*/
		number |= 1 << i
	}
	return number
}

func main() {
	number := setBit(0, 2, 1)
	fmt.Printf("%b\n", number)

	number = setBit(16, 63, 1)
	fmt.Printf("%d", number)
}
