package main
import "fmt"
func Ups() any {
	return true
	// return 1
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)

}