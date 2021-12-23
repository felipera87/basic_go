package main

import (
	"fmt"
)

// only one init per file
func init() {
	fmt.Println("this will be the first thing executed")
}
