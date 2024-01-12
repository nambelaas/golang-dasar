package main

import "fmt"

func main() {
	counter := 10

	// for counter<=10{
	// 	fmt.Println("perulangan ke", counter)
	// 	counter++
	// }

	for i := 0; i < counter; i++ {
		fmt.Println("perulangan ke", counter)

	}
	fmt.Println("perulangan ke selesai")
	
	names := []string{"Salman","Andi","Joko"}
	
	for index, name:= range names{
		fmt.Println("index",index,"=",name)
	}
	
	for _, name:= range names{
		fmt.Println(name)
	}
}
