package main

import (
	"fmt"
	"math"
)

// Структура с параметрами x, y
type Point struct {
	x float64
	y float64
}

// Конструктор для Point
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Метод, для вычисления расстояни между точками
func (p *Point) Distance(q *Point) float64 {
	return math.Sqrt(math.Pow(q.x-p.x, 2) + math.Pow(q.y-p.y, 2))
}

func task24() {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(3, 4)

	dist := p1.Distance(p2)

	fmt.Printf("Расстояние между точками %.2f\n", dist)
}
