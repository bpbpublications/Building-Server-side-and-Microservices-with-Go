package main

import "fmt"

func sum(channel chan int) int {
	s := 0

	for v := range channel {
		s = s + v
	}

	return s
}

func main() {
	channel := make(chan int, 5)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}

		close(channel)
	}()

	fmt.Println(sum(channel))
}
