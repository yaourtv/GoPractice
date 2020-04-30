package main

import (
	"fmt"
	"math"
)

type Figure interface {
	area() float64
	perimeter() float64
}

type Square struct {
	length float64
}
type Circle struct {
	radius float64
}

func (s Square) perimeter() float64 {
	return s.length * 4
}

func (s Square) area() float64 {
	return s.length * s.length
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	var s Figure = Square{}
	var c Figure = Circle{}

	fmt.Println(s.area(), s.perimeter())
	fmt.Println(c.area(), c.perimeter())

	var s1 Figure = Square{4}
	var c1 Figure = Circle{4}

	fmt.Println(s1.area(), s1.perimeter())
	fmt.Println(c1.area(), c1.perimeter())

	var s2 Figure = Square{1}
	var c2 Figure = Circle{5}

	fmt.Println(s2.area(), s2.perimeter())
	fmt.Println(c2.area(), c2.perimeter())
}