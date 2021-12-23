package main

import (
	"fmt"
	"reflect"
)

func main() {

	var array1 [5]int
	array1[0] = 2
	fmt.Println(array1)

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice)
	fmt.Println(reflect.TypeOf(slice))  // []int
	fmt.Println(reflect.TypeOf(array1)) // [5]int

	slice = append(slice, 7)
	fmt.Println(slice)

	// slice is a chunk of an array, it is a pointer to a part of an array
	slice2 := array2[1:3]
	fmt.Println(slice2) // [2 3]
	array2[1] = 9
	fmt.Println(slice2) // [9 3]

	// internal array
	// slice always points to an array, if the array is not specified then it uses the internal array

	slice3 := make([]int, 5, 7)
	fmt.Println(slice3, len(slice3), cap(slice3))
	// [0 0 0 0 0] 5 7

	slice3 = append(slice3, 1)
	slice3 = append(slice3, 1)
	fmt.Println(slice3, len(slice3), cap(slice3))
	// [0 0 0 0 0 1 1] 7 7

	// overflow slice with capacity 7
	slice3 = append(slice3, 2)
	fmt.Println(slice3, len(slice3), cap(slice3))
	// [0 0 0 0 0 1 1 2] 8 14

	// the last parameter (capacity) is optional
	slice4 := make([]int, 5)
	fmt.Println(slice4, len(slice4), cap(slice4))
	// [0 0 0 0 0] 5 5
}
