package main

import "fmt"

func execNamelessFunction() {
	result := func(someText string) string {
		fmt.Println("hello world of nameless functions")

		return fmt.Sprintf("Parameter: %s", someText)
	}("a parameter")

	fmt.Println(result)
}
