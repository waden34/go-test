package main

import "fmt"

func main() {

	// // Iterate over simple range
	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	// // Iterate over slice (array)
	// numbers := []int{1, 2, 3, 4, 5, 6}
	// for index, value := range numbers {
	// 	fmt.Printf("Index: %d, Value %v\r\n", index, value)
	// }

	// Break and continue
	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println("Odd Number:", i)
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	// rows := 5

	// // Outer loop
	// for i := 1; i <= rows; i++ {
	// 	// Inner loop for spaces before stars
	// 	for j := 1; j <= rows-i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	for k := 1; k <= 2*i-1; k++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	for i := range 10{
		fmt.Println(10-i)
	}
}
