package main

import "fmt"

func main() {
	m := make(map[int]string)

	m[1] = "Monday"
	m[3] = "Wednesday"

	value, ok := m[2]

	fmt.Println(value, ok)
}
