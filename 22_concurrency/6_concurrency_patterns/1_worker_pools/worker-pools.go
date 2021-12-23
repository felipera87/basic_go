package main

import "fmt"

// worker pools pattern inserts values in a pool which is consumed by many workers
func main() {
	tasks := make(chan int, 40)
	results := make(chan int, 40)

	// start 4 workers
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)

	// insert everything I want to process on the channel all at once
	for i := 0; i < 40; i++ {
		tasks <- i
	}
	close(tasks)

	// start collecting results produced by the workers
	for i := 0; i < 40; i++ {
		result := <-results
		fmt.Println(result)
	}
	close(results)
}

// this syntax indicates that tasks just receives data and results just sends data,
// they're not bi-direction in this scope
func worker(tasks <-chan int, results chan<- int) {
	// the worker just take a number from the pool and process it
	for number := range tasks {
		results <- fibonacci(number)
	}
}

func fibonacci(position int) int {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}
