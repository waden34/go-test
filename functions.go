package main

import "fmt"

func main() {
	//sum := add(1, 2)
	//fmt.Println(add(2, 35))
	//
	//greet := func() {
	//	fmt.Println("World")
	//}
	//fmt.Println("Hello")
	//greet()
	//operation := add
	//
	//result := operation(3, 5)
	//fmt.Println(result)

	// Passing a function as an argument
	result := applyOperation(5, 3, add)
	fmt.Println("5 + 3 = ", result)

	// returning and using a function
	multiplyBy2 := createMultiplier(2)
	fmt.Println("6 * 2 = ", multiplyBy2(6))

}

func add(a, b int) int {
	return a + b
}

// Function that takes a function as an argument
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// Function that returns a function
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
