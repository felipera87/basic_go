package main

import "fmt"

func recoverExec() {
	fmt.Println("attempt to recover")

	// the recover() will return something different of nil only if panic is called at some point
	if r := recover(); r != nil {
		fmt.Println("recover successful")
	}
}

func studentPassedExam(grade1, grade2 float32) bool {
	defer fmt.Println("Calculation process finished")
	defer recoverExec()
	fmt.Println("Initiated calculation process")

	mean := (grade1 + grade2) / 2
	if mean > 6 {
		return true
	} else if mean < 6 {

		return false
	}

	// this will exec all defers and crash the application if there's no recover function
	// this is NOT an error, it is system crash, can only be handled by recover()
	panic("MEAN IS EXACLTY 6!")
}

func main() {
	studentPassedExam(6, 6)
	// this is the awful result without recover
	/*
		Initiated calculation process
		Calculation process finished
		panic: MEAN IS EXACLTY 6!

		goroutine 1 [running]:
		main.studentPassedExam(0x40c0000040c00000, 0xc000060000)
		        /home/felipe/Documents/study/golang/basics/17_panic_recover/main.go:28 +0x1b2
		main.main()
		        /home/felipe/Documents/study/golang/basics/17_panic_recover/main.go:32 +0x35
		exit status 2
	*/

}
