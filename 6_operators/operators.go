package main

import "fmt"

func main() {
	// aritmethics
	n1 := 10 + 2
	n2 := 10 - 2
	n3 := 10 / 2
	n4 := 10 * 2
	n5 := 10 % 2
	fmt.Println(n1, n2, n3, n4, n5)

	// assign operators are = or :=

	// relational
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	// logical
	isTrue, isFalse := true, false
	fmt.Println(isTrue && isFalse)
	fmt.Println(isTrue || isFalse)
	fmt.Println(!isTrue)

	// unary
	number := 10
	number++
	number += 2
	number--
	number -= 2
	number *= 2
	number /= 2
	number %= 2
	fmt.Println(number)

}
