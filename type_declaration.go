package main

import "fmt"

func main() {
	type noKtp string

	var ktpSalman noKtp = "11111"
	var contoh string = "22222"

	var contohKtp noKtp = noKtp(contoh)

	fmt.Println(ktpSalman)
	fmt.Println(contohKtp)
}