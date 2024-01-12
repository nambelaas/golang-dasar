package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"Validation Error"}
	}

	if id != "Salman" {
		return &notFoundError{"Data Not Found"}
	}

	return nil
}

func main() {
	err := SaveData("Salman", nil)

	if err != nil {
		// terjadi error
		// if validationErr, ok := err.(*validationError); ok {
		// 	fmt.Println("Validation error:", validationErr.Error())
		// } else if notFoundErr, ok := err.(*notFoundError); ok {
		// 	fmt.Println("Not found error:", notFoundErr.Error())
		// } else {
		// 	fmt.Println("Unknown error:", err.Error())
		// }

		switch finalErr := err.(type) {
		case *validationError:
			fmt.Println("Validation error:", finalErr.Error())
		case *notFoundError:
			fmt.Println("Not found error:", finalErr.Error())
		default:
			fmt.Println("Unknown error:", err.Error())
		}
	} else {
		// sukses
		fmt.Println("Sukses")
	}
}
