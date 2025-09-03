package main

import "fmt"

func main() {
	process(10)

	process(-3)
}
func process(input int) {
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	if input < 0 {
		fmt.Println("Before panic")
		panic("input must be a non-negative number.")
	}
	fmt.Println("Processing:", input)
}
