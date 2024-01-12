package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var address Address = Address{}
	// var address *Address = &Address{}
	changeCountryToIndonesia(&address)
	// changeCountryToIndonesia(address)

	fmt.Println(address)
}
