package main

import (
	"fmt"
	"math"
)

// Point структура, хранящая кооординаты точки
type Point struct {
	x, y float64
}

// конструктор
func newPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

// расчет расстояние по формуле Пифагора
func (p1 Point) getDistance(p2 Point) float64 {
	xDist := math.Pow(p1.x-p2.x, 2)
	yDist := math.Pow(p1.y-p2.y, 2)
	return math.Sqrt(xDist + yDist)
}

func main() {
	p1 := newPoint(1, 1)
	p2 := newPoint(1, -1)

	fmt.Print(p1.getDistance(p2))
}
