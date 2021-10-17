package main

import (
	"fmt"
	"time"
)

func main() {
	red := make(chan string)
	green := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		red <- "red"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		green <- "green"
	}()

	for i := 0; i < 2; i++ {
		select {
		case color := <-red:
			fmt.Print(color + " ")
		case color := <-green:
			fmt.Println(color + " ")
		}
	}
}
