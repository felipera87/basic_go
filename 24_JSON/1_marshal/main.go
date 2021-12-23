package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   uint   `json:"age"`
}

func main() {
	// main functions of encoding is json.Marshal() which converts a map or a struct to json
	// and json.Unmarshal() which do the reverse operation (json to map or struct)

	d1 := dog{"Blica", "Shih Tzu", 8}

	jsonDog, err := json.Marshal(d1)
	if err != nil {
		log.Fatal(err)
	}

	// this return a slice of bytes: []byte or []uint8
	fmt.Println(jsonDog)
	// [123 34 110 97 109 101 34 58 34 66 108 105 99 97 34 44 34 98 114 101 101 100 34 58 34 83 104 105 104 32 84 122 117 34 44 34 97 103 101 34 58 56 125]

	// to read this slice of bytes we use bytes package
	fmt.Println(bytes.NewBuffer(jsonDog))
	// {"name":"Blica","breed":"Shih Tzu","age":8}

	d2 := map[string]string{
		"name":  "Peteca",
		"breed": "pooch",
		"age":   "7",
	}

	jsonDog2, err := json.Marshal(d2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bytes.NewBuffer(jsonDog2))
}
