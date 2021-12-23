package main

import "fmt"

type user struct {
	name string
	age  uint8
}

// this is the syntax to associate the save function to the user type
// u is a copy of the user, its like readonly, the changes made to u will be only on the scope of save()
func (u user) save() {
	fmt.Printf("%s saved!\n", u.name)
}

// u is a reference of the user, changes made on u will change its original value
func (u *user) incrementAge() {
	// no need to dereference, just use u and not *u
	u.age++
}

func main() {
	user1 := user{"felipe", 20}
	fmt.Println(user1)
	user1.save()

	user1.incrementAge()
	fmt.Println(user1)
}
