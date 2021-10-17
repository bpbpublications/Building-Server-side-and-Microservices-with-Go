package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]int)

	waitGroup := &sync.WaitGroup{}
	mapMutex := &sync.Mutex{}
	waitGroup.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			defer waitGroup.Done()
			mapMutex.Lock()
			m[i] = i
			mapMutex.Unlock()
		}(i)
	}

	waitGroup.Wait()
	fmt.Println(len(m))
}
