package main

import "fmt"

func main() {
	process(10)
}

func process(i int) {
	// arguments are evaluated in order before execution
	defer fmt.Println("Defered i value:", i)
	i++

	// Defers happen last in first out
	defer fmt.Println("This is the first deferred statement.")
	defer fmt.Println("This is the second deferred statement.")
	defer fmt.Println("This is the third deferred statement.")
	fmt.Println("This is a normally executed statement")
	fmt.Println("Value of i:", i)
}
