package main

import "fmt"

func main() {
	// fruit := "pineapple"
	// switch fruit {
	// case "apple":
	// 	fmt.Println("It's an apple")
	// case "banana":
	// 	fmt.Println("It's a banana")
	// default:
	// 	fmt.Println("It's an unknown fruit")
	// }

	// day := "Monday"
	// switch day {
	// case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
	// 	fmt.Println("It's a weekday")
	// case "Saturday", "Sunday":
	// 	fmt.Println("It's the weekend")
	// default:
	// 	fmt.Println("invalid day")
	// }

	// number := 15
	// switch {
	// case number < 10:
	// 	fmt.Println("Number is less than 10")
	// case number >= 10 && number < 20:
	// 	fmt.Println("Number is between 10 and 19")
	// default:
	// 	fmt.Println("Number is 20 or greater")
	// }

	// num := 2
	// switch {
	// case num > 1:
	// 	fmt.Println("Greater than 1")
	// 	fallthrough
	// case num == 2:
	// 	fmt.Println("num == 2")
	// default:
	// 	fmt.Println("Not 2")
	// }
	checkType(10)
	checkType(3.14)
	checkType("")
	checkType(true)

}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("It's an integer")
	case float64:
		fmt.Println("It's a float")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("It's an unknown type")
	}
}
