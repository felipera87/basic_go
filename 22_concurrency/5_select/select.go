package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			ch1 <- "channel 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			ch2 <- "channel 2"
		}
	}()

	for {
		// this way the channel1 (faster) can print 4 messages while channel2 (slower) prints 1
		// one channel doesn't wait for the other
		select {
		case messageChannel1 := <-ch1:
			fmt.Println(messageChannel1)
		case messageChannel2 := <-ch2:
			fmt.Println(messageChannel2)
		}
	}
}
