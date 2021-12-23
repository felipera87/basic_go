package main

import "fmt"

type user struct {
	name string
	age  uint8
}

type user2 struct {
	name    string
	age     uint8
	address address
}

type address struct {
	street string
	number uint8
}

func main() {
	var u1 user
	fmt.Println(u1) // print zero values of all fields

	u1.name = "felipe"
	u1.age = 20
	fmt.Println(u1)

	// declare using inference, need to pass all fields
	u2 := user{"felipe", 20}
	fmt.Println(u2)

	// declare using inference and pass only some fields
	u3 := user{name: "felipe"}
	fmt.Println(u3)

	// nested structs
	address := address{"street", 10}
	u4 := user2{"felipe", 20, address}
	fmt.Println(u4)

}
