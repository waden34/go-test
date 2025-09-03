package main

import "fmt"

func main() {
	//func functionName(param1 type1, param2 type2, param3 ...type3) returnType

	fmt.Println("The sum of these numbers is: ", sum(1, 2, 3, 5345, 364, 47))
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
