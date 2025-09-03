package main

import "fmt"

func main() {
	//func functionName(param1 type1, param2 type2, param3 ...type3) returnType

	fmt.Println("The sum of these numbers is: ", sum(1, 2, 3, 5345, 364, 47))

	numbers := []int{1, 2, 3, 4, 5, 9}
	fmt.Println("The sum of the slice is: ", sum(numbers...))
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
