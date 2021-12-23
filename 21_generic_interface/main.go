package main

import "fmt"

func generic(any interface{}) {
	fmt.Println(any)
}

func main() {
	generic("test")
	generic(1)
	generic(1.5)
	generic(true)

	genericMap := map[interface{}]interface{}{
		1:      "string",
		"name": 2,
		true:   5.8,
	}
	fmt.Println(genericMap)
}
