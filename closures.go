package main

import "fmt"

func main() {
	//sequence := adder()
	//fmt.Println(sequence())
	//fmt.Println(sequence())
	//fmt.Println(sequence())
	//fmt.Println(sequence())
	//fmt.Println(sequence())
	//
	//sequence2 := adder()
	//fmt.Println(sequence2())
	//fmt.Println(sequence())

	subtracter := func() func(int) int {
		countdown := 99
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	// Using the closure subtracter
	fmt.Println(subtracter(1))
	fmt.Println(subtracter(2))
	fmt.Println(subtracter(3))
	fmt.Println(subtracter(4))
	fmt.Println(subtracter(5))
	fmt.Println(subtracter(6))
}

func adder() func() int {
	i := 0
	fmt.Println("previous value of i:", i)
	return func() int { // Anonymous function that is set to the sequence variable above. Lines above are only called in the initial adder() call
		i++
		fmt.Println("added 1 to i:", i)
		return i
	}
}
