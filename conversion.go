package main

import "fmt"

func main() {
	// Ada kondisi number overflow jika angka didalam variable int melebihi kapasitasnya maka akan menampilkan nilai terbawahnya dari variable pertamanya
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)


	var name = "Salman Seif"
	var e uint8 = name[0]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}