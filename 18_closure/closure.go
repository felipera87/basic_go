package main

import "fmt"

func closureTest() func() {
	someText := "inside closureTest scope"

	someFunction := func() {
		fmt.Println(someText)
	}

	return someFunction
}

func main() {
	someText := "outside closureTest scope"
	fmt.Println(someText)

	f := closureTest()
	f() // inside closureTest scope
}
