package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	var salman Customer

	salman.Name = "Salman Seif"
	salman.Address = "Indonesia"
	salman.Age = 23

	fmt.Println(salman)

	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     25,
	}
	fmt.Println(joko)

	budi := Customer{"Budi",
		"Indonesia", 27}

	fmt.Println(budi)

	budi.sayHello("Agus")
}
