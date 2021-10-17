package main

import "fmt"

func inc(a int) int {
	fmt.Println("Increment:", a)
	return a + 1
}

func main() {
	i := 10

	switch {
	case inc(i) < 5:
		fmt.Println("Lower than five")
	case inc(i) > 5:
		fmt.Println("Greater than five")
	default:
		fmt.Println("Five")
	}
}
