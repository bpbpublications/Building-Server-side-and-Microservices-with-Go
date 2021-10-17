package main

import "fmt"

func main() {
	fmt.Print("a")
	defer fmt.Print("b")
	fmt.Print("c")
}
