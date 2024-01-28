package main

import "fmt"

func getIntersection(set1, set2 map[int]bool) map[int]bool {
	res := make(map[int]bool)

	/* итерируемся по ключам множества set2, если ключ содержится в множестве set1,
	   добавляем значение ключа в результирующее множество*/
	for key := range set2 {
		if set1[key] {
			res[key] = true
		}
	}
	return res
}

func main() {
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)

	for i := 1; i <= 10; i++ {
		set1[i] = true
		set2[i] = true
	}

	for i := 11; i < 20; i++ {
		set2[i] = true
	}

	fmt.Print(getIntersection(set1, set2))
}
