package main

import (
	"errors"
	"fmt"
)

func main() {
	var (
		num1 int8
		num2 int16
		num3 int32
		num4 int64
	)
	fmt.Println(num1, num2, num3, num4)

	// normal int uses pc architecture, could be 32 or 64
	var normalInt int = 42
	fmt.Println(normalInt)

	// same for int applies to unsigned int
	var unsignedNumber uint = 10
	fmt.Println(unsignedNumber)

	// special name for int32, used to work with ascii characters
	var thisIsInt32 rune = 64
	fmt.Println(thisIsInt32)

	// special name for uint8
	var thisIsInt8 byte = 255
	fmt.Println(thisIsInt8)

	var (
		real1 float32 = 0.32
		real2 float64 = 0.64
	)
	fmt.Println(real1, real2)

	// just "float" doesn't exist, only when using inference.
	// in this case it is based on the pc architecture, same as int
	real3 := 32.64
	fmt.Println(real3)

	var str1 string = "this is a string"
	str2 := "also a string, but using type inference"
	fmt.Println(str1, str2)

	// char type doesn't exist, it's actually a int and points to ascii table
	char := 'B'
	fmt.Println(char)

	var boolean bool = true
	fmt.Println(boolean)

	// error is a type in go, and it's zero value is nil
	var err error
	fmt.Println(err)

	var err2 error = errors.New("Custom error using go internal package")
	fmt.Println(err2)
}
