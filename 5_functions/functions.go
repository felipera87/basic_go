package main

import "fmt"

func sum(n1, n2 int8) int8 {
	return n1 + n2
}

// multiple returns
func manyCalcs(n1, n2 int) (int, int) {
	return n1 + n2, n1 - n2
}

func main() {
	sumResult := sum(1, 2)
	fmt.Println(sumResult)

	var f = func(txt string) {
		fmt.Println("function in a variable", txt)
	}
	f("custom text")

	// using many return functions
	sumResult2, subResult := manyCalcs(20, 10)
	fmt.Println(sumResult2, subResult)

	// to omit any result
	_, subResult2 := manyCalcs(20, 10)
	fmt.Println(subResult2)
}
