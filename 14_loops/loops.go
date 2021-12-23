package main

import (
	"fmt"
	"time"
)

func main() {
	// loops in go only have FOR statement

	// while like statement
	i := 0

	for i < 5 {
		i++
		fmt.Println("i", i)
	}

	for j := 0; j < 5; j++ {
		fmt.Println("j", j)
	}

	names := [3]string{"felipe", "eliria", "bob"}

	for index, name := range names {
		fmt.Println(index, name)
	}

	for index := range names {
		fmt.Println(index)
	}

	for _, name := range names {
		fmt.Println(name)
	}

	for index, character := range "some text" {
		fmt.Println(index, character, string(character))
		// character will be the ascii number
	}

	user := map[string]string{
		"name":     "felipe",
		"lastname": "reis",
	}

	for key, value := range user {
		fmt.Println(key, value)
	}

	k := 0
	// infinite loop
	for {
		fmt.Println("step", k)
		time.Sleep(time.Second)
		k++

		if k == 3 {
			break
		}
	}
}
