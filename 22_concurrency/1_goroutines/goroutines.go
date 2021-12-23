package main

import (
	"fmt"
	"time"
)

// paralelism is when code runs in different cores at the same time
// concurrency is when code looks like it's running at the same time,
// but not necessarily. It runs independently but can use same hardware resources

func main() {
	// if a function is called with "go", it is a goroutine and runs in concurrency
	go write("hello world")

	// if put "go" here it doesn't print anything
	write("hello world2")
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
