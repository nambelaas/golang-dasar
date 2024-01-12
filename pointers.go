package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Brebes", "Jawa Tengah", "Indonesia"}
	// address2 := address1

	// address2.City = "Semarang"

	// fmt.Println(address1)
	// fmt.Println(address2)

	address1 := Address{"Brebes", "Jawa Tengah", "Indonesia"}
	// var address1 Address = Address{"Brebes", "Jawa Tengah", "Indonesia"}
	address2 := &address1 // pointer
	// var address2 *Address = &address1 // pointer

	address2.City = "Semarang"

	fmt.Println(address1)
	fmt.Println(address2)
}
