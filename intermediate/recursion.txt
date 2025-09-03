package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
	fmt.Println(factorial(10))

	fmt.Println(sumOfDigits(124))
	fmt.Println(sumOfDigits(14))
	fmt.Println(sumOfDigits(4))

	for i := range 50 {
		fmt.Println(fibonacci(i))
	}
}

func factorial(n int) int {
	// Base case factorial of 0 is 1
	if n == 0 {
		return 1
	}
	// Recursive case: factorial of n is n * factorial(n-1)
	return n * factorial(n-1)
	// n * (n - 1) * factorial (n-2)... factorial(0)

}

func sumOfDigits(n int) int {
	// Base case
	if n < 10 {
		return n
	}
	return n%10 + sumOfDigits(n/10)
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
