package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	err := checkmail.ValidateFormat("test@test.com")
	fmt.Println(err)
	err = checkmail.ValidateFormat("aaaa")
	fmt.Println(err)
}
