package main

import "fmt"

func main() {
	var var1 string = "var1" // define type explicitly
	var2 := "var2"           // inference
	fmt.Println(var1, var2)

	// multiple variables at once
	var (
		var3 string = "var3"
		var4 string = "var4"
	)
	fmt.Println(var3, var4)

	// multiple variables at once using inference
	var5, var6 := "var5", "var6"
	fmt.Println(var5, var6)

	// constants
	const const1 string = "const1"
	fmt.Println(const1)
}
