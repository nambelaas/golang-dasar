package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	// teknik recover yg benar adalah diletakkan pada fungsi yg di defer
	message := recover()
	fmt.Println("Terjadi Error", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Ups Error")
	}

	// teknik recover yg salah
	// message:=recover()
	// fmt.Println("Terjadi Error",message)
}

func main() {
	runApp(false)
	fmt.Println("Salman Seif")
}
