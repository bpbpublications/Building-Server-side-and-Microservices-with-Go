package main

import "fmt"

func main() {
	arr := [5]int{5, 7, 4, 9, 3}
	sum := 0

	for i, v := range arr {
		sum = sum + v
		fmt.Println("index: ", i, "value: ", v)
	}

	fmt.Println("sum: ", sum)
}
