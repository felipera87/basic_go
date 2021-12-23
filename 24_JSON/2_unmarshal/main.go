package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// to omit a json key it's possible to do this: `json:"-"`
// the data in the struct will exist but will not be exported when using unmarshal or marshal
type dog struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   uint   `json:"age"`
}

func main() {
	// json to struct
	jsonDog := `{"name":"Blica","breed":"Shih Tzu","age":8}`

	var d1 dog

	// Unmarshal must receive a slice of bytes. It's not necessary to use the
	// byte package, we can cast the string to a slice of bytes
	if err := json.Unmarshal([]byte(jsonDog), &d1); err != nil {
		log.Fatal(err)
	}

	fmt.Println(d1)

	// json to map
	jsonDog2 := `{"age":"7","breed":"pooch","name":"Peteca"}`

	d2 := make(map[string]string)

	if err := json.Unmarshal([]byte(jsonDog2), &d2); err != nil {
		log.Fatal(err)
	}

	fmt.Println(d2)
}
