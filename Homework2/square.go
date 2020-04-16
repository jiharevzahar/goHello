package main

import "fmt"

//Point : coordinates
type Point struct {
	x, y int
}

//Square : starting point and side size
type Square struct {
	start Point
	a     uint
}

//Perimeter : perimeter of square
func (s Square) Perimeter() uint {
	return s.a * 4
}

//Area : area of square
func (s Square) Area() uint {
	return s.a * s.a
}

//End : finds end point
func (s Square) End() (int, int) {
	//this is the first solution i found, maybe i need to return Point
	return int(s.start.x + int(s.a)), int(s.start.y - int(s.a))
}

func main() {

	s := Square{Point{1, 1}, 5}
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())
	fmt.Println(s.End())
}
