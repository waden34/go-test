package main

import "fmt"

func main() {
	// var ptr *int

	var a int = 10

	var ptr *int

	ptr = &a // Referencing

	fmt.Println(a)
	fmt.Println(ptr) // memory address where the a variable is stored
	//
	//if ptr == nil {
	//	fmt.Println("ptr is nil")
	//} else {
	//	fmt.Println(*ptr) // dereferencing a pointer
	//}

	modifyValue(ptr)
	fmt.Println(a)
	modifyVariable(a)
	fmt.Println(a)
}

func modifyValue(ptr *int) {
	*ptr++ // defreference and increment the pointer. Updates the underlying variable value. Non-pointer params would be copies of variables and changes wouldn't persist in outer scope
}

func modifyVariable(i int) {
	i++
}
