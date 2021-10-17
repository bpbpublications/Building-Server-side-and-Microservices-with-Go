package main

import "fmt"

func inc(i int) int {
	return i + 1
}

func incP(pi *int) {
	*pi = *pi + 1
}

func main() {
	x := 5
	inc(x)
	fmt.Println(x)

	incP(&x)
	fmt.Println(x)
}
