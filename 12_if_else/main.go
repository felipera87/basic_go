package main

import "fmt"

func main() {
	var1 := 10
	if var1 > 10 {
		fmt.Println("bigger than 10")
	} else {
		fmt.Println("not bigger than 10")
	}

	// initialize if statement with a variable
	// this variable is only available within the if statement scope
	if var2 := var1; var2 > 0 {
		fmt.Println("var2 bigger than 0")
	} else if var2 < 0 {
		fmt.Println("var2 less than 0")
	}
}
