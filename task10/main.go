package main

import (
	"fmt"
	"math"
)

func main() {
	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	mp := make(map[int][]float64)

	for _, t := range temp {
		// в качестве ключа используем саму температуру t, из которой вычтена ее дробная часть
		key := int(t - math.Mod(t, 10))
		// если ключа еще нет в mp, инициализируем срез
		if _, ok := mp[key]; !ok {
			mp[key] = []float64{}
		}
		mp[key] = append(mp[key], t)
	}

	for key, slice := range mp {
		fmt.Printf("key: %3d - slice %v\n", key, slice)
	}
}
