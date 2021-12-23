package main

import (
	"fmt"
	"time"
)

func main() {
	channel := write("hello world")

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

// the channel returned just receives data
func write(text string) <-chan string {
	ch := make(chan string)

	// this will run forever, as soon as 1 iteration completes it sends one more value to the channel
	// if the channel is full, it waits until the main thread pulls data from it
	go func() {
		for {
			ch <- fmt.Sprintf("text: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return ch
}
