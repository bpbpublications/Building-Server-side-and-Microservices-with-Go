package main

import "fmt"

var globalA int
var globalB = 1

func main() {
	var localA int
	var localB, localC int = 2, 3

	fmt.Println(globalA, globalB, localA, localB, localC)
}
