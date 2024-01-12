package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}
func main() {
	amount := sumAll(10, 10, 10, 10, 10, 10)
	fmt.Println(amount)
	
	numbers:=[]int{10,10,10}
	// fmt.Println(sumAll(numbers)) error karena terbaca 1 data
	fmt.Println(sumAll(numbers...))

}
