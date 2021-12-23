package main

import "fmt"

func dynamicParamaterSum(numbers ...int) int {
	// numbers is a slice
	fmt.Println(numbers) // [1 2 3 4 5 6]

	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

// there can be only one dynamic parameter, and it must be the last one
func dynamicParamaterWriteGarbage(someText string, numbers ...int) {
	var numberString string
	for _, number := range numbers {
		numberString += fmt.Sprint(number)
	}

	fmt.Println(someText, numberString)
}
