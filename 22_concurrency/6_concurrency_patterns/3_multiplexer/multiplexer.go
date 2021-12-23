package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := multiplex(write("channel 1"), write("channel 2"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

// this function will merge 2 channels in one
func multiplex(inChannel1, inChannel2 <-chan string) <-chan string {
	outChannel := make(chan string)

	go func() {
		for {
			select {
			case msg := <-inChannel1:
				outChannel <- msg
			case msg := <-inChannel2:
				outChannel <- msg
			}
		}
	}()

	return outChannel
}

func write(text string) <-chan string {
	ch := make(chan string)

	go func() {
		for {
			ch <- fmt.Sprintf("text: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return ch
}
