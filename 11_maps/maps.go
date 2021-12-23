package main

import "fmt"

func main() {
	// map[type_of_key]type_of_value
	user := map[string]string{
		"name":      "felipe",
		"last_name": "reis",
	}

	fmt.Println(user)
	fmt.Println(user["name"])

	// nested maps
	user2 := map[string]map[string]string{
		"name": {
			"first": "felipe",
			"last":  "reis",
		},
		"address": {
			"street": "some street",
			"number": "some number",
		},
	}

	fmt.Println(user2)

	// delete a key
	delete(user2, "address")
	fmt.Println(user2)

	// add a key
	user2["address2"] = map[string]string{
		"street": "some street",
	}
	fmt.Println(user2)
}
