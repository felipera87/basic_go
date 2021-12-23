package main

import "fmt"

func function1() {
	fmt.Println("exec function 1")
}

func function2() {
	fmt.Println("exec function 2")
}

func studentPassedExam(grade1, grade2 float32) bool {
	defer fmt.Println("Calculation process finished")
	fmt.Println("Initiated calculation process")

	mean := (grade1 + grade2) / 2
	if mean >= 6 {
		return true
	}
	return false
}

func main() {
	// this will exec function2 first, then function1
	defer function1()
	function2()

	fmt.Println(studentPassedExam(6, 5))
}
