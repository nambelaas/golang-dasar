package main

import "fmt"

func main() {
	var a = 1
	var b = 5
	var c = 7
	var d = a+b*c

	fmt.Println(d)

	// augmented assignment
	var i = 10
	i+=10
	fmt.Println(i)
	i+=5
	fmt.Println(i)

	// unary operator
	var j = 1
	j++
	fmt.Println(j)
	j++
	fmt.Println(j)
	j--
	fmt.Println(j)
}