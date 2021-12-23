package main

import "fmt"

func main() {
	var var1 int = 10
	var var2 int = var1

	fmt.Println(var1, var2) // 10 10
	var1++
	fmt.Println(var1, var2) // 11 10

	var var3 int
	var pointer *int

	fmt.Println(var3, pointer) // 0 <nil>

	var3 = 100
	pointer = &var3             // put address of var3 in pointer
	fmt.Println(var3, pointer)  // 100 0xc000084038
	fmt.Println(var3, *pointer) // 100 100

	var3++
	fmt.Println(var3, *pointer) // 101 101
}
