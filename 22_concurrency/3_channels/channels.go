package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go write("hello world", channel)
	for {
		// this instruction halts the execution and waits for something to get out on this channel
		message, isChannelOpen := <-channel
		if !isChannelOpen {
			break
		}
		fmt.Println(message)
	}

	// this is same effect as the code above using range
	channel2 := make(chan string)
	go write("hello world2", channel2)
	for message := range channel2 {
		fmt.Println(message)
	}

	fmt.Println("end")
}

func write(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(time.Second)
	}

	close(channel)
}
