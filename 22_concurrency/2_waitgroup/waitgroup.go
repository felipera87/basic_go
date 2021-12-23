package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	// tell that 2 goroutines will be executed
	waitGroup.Add(2)

	go func() {
		write("hello world")
		waitGroup.Done()
	}()

	go func() {
		write("hello world2")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
