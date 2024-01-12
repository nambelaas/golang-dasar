package main

import "fmt"

func main() {
	name := "Joko"

	if name == "Budi" {
		fmt.Println("Hello", name)
	} else if name == "Joko" {
		fmt.Println("Hello", name)
	} else {
		fmt.Println("Hello Guest")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
