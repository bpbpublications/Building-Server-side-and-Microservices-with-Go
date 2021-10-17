package main

import "fmt"

type Rectangle struct {
	a, b int
}

func (r *Rectangle) Area() int {
	return r.a * r.b
}

func main() {
	r := Rectangle{2, 3}
	fmt.Println(r.Area())
}
