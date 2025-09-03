package main

import "fmt"

func main() {

	//// Printing Functions
	//fmt.Print("Hello ")
	//fmt.Print("World!")
	//fmt.Print(12, 456)
	//
	//fmt.Println("Hello ")
	//fmt.Println("World!")
	//fmt.Println(12, 456)
	//
	//name := "John"
	//age := 25
	//fmt.Printf("Name: %s, Age: %d\n", name, age)
	//fmt.Printf("Binary: %b, Hex: %x\n", age, age)
	//
	//// Formatting Functions
	//s := fmt.Sprint("Hello", "World!", 123, 456)
	//fmt.Print(s)
	//fmt.Print(s)
	//
	//s = fmt.Sprintln("Hello", "World!", 123, 456)
	//fmt.Print(s)
	//fmt.Print(s)
	//
	//s = fmt.Sprintf("Name: %s, Age: %d", name, age)
	//fmt.Println(s)

	// Scanning Functions (Getting input from console)
	//var name string
	//var age int
	//fmt.Println("Enter your name and age:") // Arguments are space separated. Can't use " to enter as a single
	////_, err := fmt.Scan(&name, &age)
	////if err != nil {
	////	return
	////}
	////fmt.Scanln(&name, &age)
	//fmt.Scanf("%s %d", &name, &age)
	//fmt.Printf("Hello %s, your age is %d\n", name, age)

	// Error Formatting Functions

	err := checkAge(15)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("age %d is less than 18", age)
	}
	return nil
}
