package main

import (
	"fmt"
	"math"
)

type Figure interface {
	Area() float64
	Perimeter() float64
}

type Square struct {
	side uint
}

func (s Square) Area() float64 {
	return float64(s.side * s.side)
}

func (s Square) Perimeter() float64 {
	return float64(4 * s.side)
}

type Circle struct {
	radius uint
}

func (c Circle) Area() float64 {
	return (float64(math.Pi) * float64(c.radius) * float64(c.radius))
}

func (c Circle) Perimeter() float64 {
	return float64(4 * c.radius)
}

func main() {
	var s Figure = Square{5}
	var c Figure = Circle{5}

	fmt.Println(s.Area(), s.Perimeter())
	fmt.Println(c.Area(), c.Perimeter())
}
