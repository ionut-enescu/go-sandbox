package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	t := triangle{15, 20}
	fmt.Print("Area for a triangle of height ", t.height, " and base of ", t.base, " is: ")
	printArea(t)

	s := square{25}
	fmt.Print("Area for a square of side length of ", s.sideLength, " is: ")
	printArea(s)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
