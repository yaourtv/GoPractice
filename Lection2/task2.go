package main

import "fmt"

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

// We don't want to change start point. Are we?
func (s Square) End() (int, int) {
	return s.start.x + int(s.a), s.start.y + int(s.a)
}

func (s Square) Perimeter() uint {
	return s.a * 4
}

func (s Square) Area() uint {
	return s.a * s.a
}

func main() {
	s := Square{Point{1, 1}, 5}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())
}
