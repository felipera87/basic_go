package main

import (
	"fmt"
	"intro/address"
)

func main() {
	addressType := address.GetType("Avenida Paulista")
	fmt.Println(addressType)
}
