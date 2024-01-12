package helper

import "fmt"

// jika variable/func diawali huruf besar maka dapat diakses dari luar, begitu pula sebaliknya
var version = "1.0.0" // tidak bisa diakses dari luar
var Application = "golang"

func sayGoodbye(name string) string {
	return "Goodbye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}

func Contoh() {
	sayGoodbye("Salman")
	fmt.Println(version)
}
