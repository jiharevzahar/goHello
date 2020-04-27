package main

import (
	"fmt"
	"math"
	"errors"
)

type Figure interface {
	area() (float64,error)
	perimeter() (float64,error)
}

type Square struct {
	side int
}

func (s Square) area() (float64,error) {
	if s.side < 0 {
		return 0,errors.New("number is less than 0")
	}
	return float64(s.side * s.side),nil
}

func (s Square) perimeter() (float64,error) {
	if s.side < 0 {
		return 0,errors.New("number is less than 0")
	}
	return float64(4 * s.side),nil
}

type Circle struct {
	radius int
}

func (c Circle) area() (float64,error) {
	if c.radius < 0{
		return 0,errors.New("number is less than 0")
	}
	return (float64(math.Pi) * float64(c.radius) * float64(c.radius)),nil
}

func (c Circle) perimeter() (float64,error) {
	if c.radius < 0{
		return 0,errors.New("number is less than 0")
	}
	return float64(4 * c.radius),nil
}

func main() {
	var s Figure = Square{-3}
	var c Figure = Circle{5}

	sArea,err:=s.area()
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(sArea)
	}

	sPerim,err:=s.perimeter()
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(sPerim)
	}

	cArea,err:=c.area()
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(cArea)
	}

	cPerim,err:=c.perimeter()
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(cPerim)
	}
}
