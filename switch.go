package main

import "fmt"

func main() {
	name := "Andi"

	switch name {
	case "Salman":
		fmt.Println("Hello", name)
	case "Joko":
		fmt.Println("Hello", name)
	default:
		fmt.Println("Hello Guest")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	default:
		fmt.Println("Nama sudah benar")
	}

	length := len(name)
	switch  {
	case length > 5:
		fmt.Println("Nama terlalu panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
