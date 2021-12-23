package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	// could be person person, but to access values need to be s.person.course
	person
	grade string
}

func main() {

	p1 := person{"felipe", 20}
	fmt.Println(p1)

	s1 := student{p1, "A"}
	fmt.Println(s1.name, s1.age, s1.grade)
	fmt.Println("s1.person.name also works", s1.person.name)
}
