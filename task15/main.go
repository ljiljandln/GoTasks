package main

func createHugeString(size uint64) string {
	return string(make([]rune, size))
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	// чтобы учитывались руны, а не байты, нужно сделать дополнительные преобразования
	// сначала получаем из строки срез рун из строки v
	slice := []rune(v)
	// обрезаем slice и преобразуем в строку
	justString = string(slice[:100])
}

func main() {
	someFunc()
}
