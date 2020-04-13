package main

import "fmt"

// Structure describes square start point (x, y)
type Point struct {
	x, y int
}

// Structure describes square with its start point and length
type Square struct {
	start Point
	// Square length value
	a     uint
}

// Method takes length and returns square end point
func (s Square) End() (int, int) {
	return s.start.x + int(s.a), s.start.y + int(s.a)
}

// Method takes length and returns square perimeter
func (s Square) Perimeter() uint {
	return s.a * 4
}

// Method takes length and returns square area
func (s Square) Area() uint {
	return s.a * s.a
}

func main() {
	s := Square{Point{1, 1}, 5}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())
}
