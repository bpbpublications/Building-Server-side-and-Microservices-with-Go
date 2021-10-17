package main

import "fmt"

func main() {
	i := 2

	switch i {
	case 1:
		fmt.Print("One")
	case 2:
		fmt.Print("Two")
	case 3:
		fmt.Print("Three")
	default:
		fmt.Println("Zero")
	}

	fmt.Print("End")

}
