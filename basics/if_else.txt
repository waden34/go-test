package main

import "fmt"

func main() {
	// age := 25
	// if age >= 18 {
	// 	fmt.Println("You are an adult")
	// // }

	// temperature := 25
	// if temperature >= 30 {
	// 	fmt.Println("It's hot outside")
	// } else {
	// 	fmt.Println("It's cool outside")
	// }

	// score := 95
	// if score >= 90 {
	// 	fmt.Println("Grade A")
	// } else if score >= 80 {
	// 	fmt.Println("Grade B")
	// } else if score >= 70 {
	// 	fmt.Println("Grade C")
	// } else if score >= 60 {
	// 	fmt.Println("Grade D")
	// } else {
	// 	fmt.Println("You Fail")
	// }
	number := 18
	if number%2 == 0 {
		if number%3 == 0 {
			fmt.Println("Number is divisible by 2 and 3")
		} else {
			fmt.Println("Number is only divisible by 2")
		}
	} else {
		fmt.Println("Number is not divisible by 2")
	}

}
