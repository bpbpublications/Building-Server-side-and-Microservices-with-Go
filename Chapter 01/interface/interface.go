package main

import "fmt"

type Printer interface {
	Print()
}

type Point struct {
	x, y int
}

func (p Point) Print() {
	fmt.Printf("(%v, %v)", p.x, p.y)
}

func main() {
	p := Point{x: 4, y: 6}
	p.Print()

	var i Printer
	i.Print()
}
