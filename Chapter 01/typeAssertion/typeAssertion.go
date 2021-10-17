package main

import "fmt"

type Rectangle struct {
	a, b int
}

type Triangle struct {
	a, h int
}

func area(i interface{}) int {
	if r, ok := i.(Rectangle); ok {
		return r.a * r.b
	}

	if t, ok := i.(Triangle); ok {
		return t.a * t.h / 2
	}

	return 0
}

func main() {
	r := Rectangle{2, 4}
	t := Triangle{5, 4}
	i := 5

	fmt.Println(area(r))
	fmt.Println(area(t))
	fmt.Println(area(i))
}
