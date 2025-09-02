package main

import "fmt"

func main() {
	//var numbers [5]int
	//fmt.Println(numbers)
	//
	//numbers[4] = 20
	//fmt.Println(numbers)
	//
	//numbers[0] = 9
	//fmt.Println(numbers)
	//
	//fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}
	//fmt.Println("Fruits array:", fruits)
	//
	//fmt.Println("Third element:", fruits[2])
	//
	//originalArray := [3]int{1, 2, 3}
	//copiedArray := originalArray
	//copiedArray[0] = 100
	//fmt.Println("Original:", originalArray)
	//fmt.Println("Copied:", copiedArray)
	//
	//for i := 0; i < len(numbers); i++ {
	//	fmt.Println("Element at index", i, ":", numbers[i])
	//}
	//
	//for i, v := range numbers {
	//	fmt.Printf("Index %d, Value %d\r\n", i, v)
	//}
	//// _ is bland identifier used to store unused variable
	//a, _ := someFunction()
	//fmt.Println(a)
	//// fmt.Prinln(b)
	//fmt.Println("The length of numbers array is:", len(numbers))
	//
	//// Compare arrays
	//array1 := [3]int{1, 2, 3}
	//array2 := [3]int{10, 2, 3}
	//fmt.Println("Array 1 is equal to Array 2:", array1 == array2)
	//
	//var matrix [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	//fmt.Println(matrix)
	originalArray := [3]int{1, 2, 3}
	var copiedArray *[3]int = &originalArray
	copiedArray[0] = 100
	fmt.Println("Original:", originalArray)
	fmt.Println("Copied:", *copiedArray)
}

func someFunction() (int, int) {
	return 1, 2
}
