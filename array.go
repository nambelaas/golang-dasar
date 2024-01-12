package main

import "fmt"

func main() {
	var nama [2] string
	nama[0] = "Salman"
	nama[1] = "Seif"

	fmt.Println(nama[0])
	fmt.Println(nama[1])
	
	// jika isi array tidak ingin di set diawal maka di set menjadi [...] agar ditentukan otomatis dan isinya wajib di set juga
	var values = [...]int{
		1,2,3,
	}
	
	fmt.Println(values)
	fmt.Println(len(values))
	fmt.Println(values)
}
