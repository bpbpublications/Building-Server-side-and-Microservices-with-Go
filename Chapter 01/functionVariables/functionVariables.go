package main

import "fmt"

func calc(param int) func(int, int) int {
	if param >= 0 {
		return func(a, b int) int {
			return a + b
		}
	}

	return func(a, b int) int {
		return a - b
	}
}

func main() {
	add := calc(5)
	sub := calc(-5)

	fmt.Println(add(3, 2), sub(3, 2))
}
