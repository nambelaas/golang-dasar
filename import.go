package main
import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result:= helper.SayHello("Salman")
	fmt.Println(result)
	bye:= helper.sayGoodbye("Salman")
	fmt.Println(bye)
}