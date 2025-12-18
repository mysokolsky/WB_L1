// L1.24
// Расстояние между точками
// Разработать программу нахождения расстояния между
// двумя точками на плоскости.

// Точки представлены в виде структуры Point с инкапсулированными (приватными)
// полями x, y (типа float64) и конструктором.
// Расстояние рассчитывается по формуле между координатами двух точек.

// Подсказка: используйте функцию-конструктор NewPoint(x, y),
// Point и метод Distance(other Point) float64.

// Дедлайн — 10 дек, 02:59

// Решение:

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (p *Point) Distance(p2 *Point) float64 {
	return math.Pow((math.Pow((p2.x-p.x), 2) + math.Pow((p2.y-p.y), 2)), 1/2.0)
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func main() {
	a := NewPoint(0, 0)
	b := NewPoint(-1, 1)
	fmt.Printf("%.5f", b.Distance(a))
}
