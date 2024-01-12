package main

import (
	"fmt"
)

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging()
	fmt.Println("Sedang menjalankan app")
}

func main() {
	runApplication()
}
