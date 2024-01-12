package main

import "fmt"

func getCompleteName() (firstName, lastName string) {
	firstName = "Salman"
	lastName = "Seif"

	return firstName, lastName
}

func main() {
	firstName, _ := getCompleteName()

	fmt.Println(firstName)
}
