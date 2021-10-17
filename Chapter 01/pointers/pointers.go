package main

import "fmt"

func main() {
	var i int = 8
	var p *int

	p = &i
	*p = 5

	fmt.Println(i)
}
