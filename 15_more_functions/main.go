package main

import "fmt"

func main() {
	result1, result2 := namedReturnCalc(5, 6)
	fmt.Println(result1, result2)

	fmt.Println(dynamicParamaterSum(1, 2, 3, 4, 5, 6))
	dynamicParamaterWriteGarbage("initial text", 1, 2, 3, 4, 5)

	execNamelessFunction()

	position := uint(10)
	fmt.Println(fibonacci(position))

	number := 10
	pointerSwapMinusSign(&number)
	fmt.Println(number)
}
