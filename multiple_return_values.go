package main

import (
	"errors"
	"fmt"
)

func main() {
	quotient, remainder := divide(10, 3)
	fmt.Printf("Quotient: %d, Remainder %d\n", quotient, remainder)

	result, err := compare(3, 3)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Result: %v\n", result)
}
func divide(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if b > a {
		return "b is greater than a", nil
	} else {
		return "", errors.New("unable to compare which number is greater")
	}
}
