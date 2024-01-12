package main

import "fmt"

func main() {
	var data map[string]string = map[string]string{
		"firstName" : "Salman",
		"lastName": "Seif", 
	}

	fmt.Println(data)

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Salman"
	book["ups"] = "Salah"

	fmt.Println(book)
	
	delete(book, "ups")
	fmt.Println(book)
}