package main

import "fmt"

func main() {
	/**
		- jika variable tidak langsung diisi maka perlu declare tipe datanya
		- var bisa dihilangkan jika variable langsung diisi dan menggunakan kata kunci :=
	**/
	// var name string
	// var name = "Salman"
	name := "Salman"
	fmt.Println("Name = ", name)

	name = "Salman Seif"
	fmt.Println("Name = ", name)

	var(
		firstName = "Salman"
		lastName = "Seif"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}