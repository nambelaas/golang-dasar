package main

import "fmt"

func main() {
	name := [...]string{
		"Salman",
		"Seif",
		"Joko",
		"Budi",
		"Andi",
		"Nugraha",
	}

	slice1 := name[3:6]
	slice2 := name[:3]
	slice3 := name[3:]
	slice4 := name[:]

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daySlice1 := days[5:]
	fmt.Println(daySlice1) // Sabtu & Minggu
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Salman"
	newSlice[1] = "Seif"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// menggunakan append untuk menambahkan array jika capacitynya masih ada
	newSlice2 := append(newSlice, "Salman")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "agus"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)
	
	iniArray := [...]int{1,2,3,4,5,}
	iniSlice := []int{1,2,3,4,5,}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
