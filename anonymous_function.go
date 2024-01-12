package main

import "fmt"

type Blacklist func(string) bool

func RegisterUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}
	RegisterUser("Salman", blacklist)
	RegisterUser("Salman", func(name string) bool {
		return name == "anjing"
	})
}
