package main

import (
	"fmt"
)

func main() {
	// a buffer is used to block the application only when buffer is full,
	// in this case with 2 elements on the channel

	// second parameter is a buffer
	channel := make(chan string, 2)

	channel <- "hello world"
	channel <- "hello world2"
	// if one more thing is sent to channel it deadlocks because overflow the buffer size

	message := <-channel
	message2 := <-channel
	// if one more thing is pulled from the channel it deadlocks because nothing
	// is sending more things to the channel

	fmt.Println(message)
	fmt.Println(message2)
}
