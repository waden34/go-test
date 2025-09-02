package main

import (
	"fmt"
	"slices"
)

func main() {
	//var numbers []int
	//var numbers1 = []int{1, 2, 3, 4, 5}
	//numbers2 := []int{1, 2, 3, 4, 5, 6}
	//slice := make([]int, 5)
	a := [5]int{1, 2, 3, 4, 5}
	slice1 := a[1:4]

	fmt.Println(slice1)

	slice1 = append(slice1, 6, 7)
	fmt.Println("Slice1", slice1)

	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice1)

	fmt.Println("SliceCopy", sliceCopy)

	var nilSlice []int
	fmt.Println(nilSlice)

	for i, v := range slice1 {
		fmt.Println(i, v)
	}

	fmt.Println(slice1[3])

	//slice1[3] = 50
	//fmt.Println(slice1[3])

	if slices.Equal(slice1, sliceCopy) {
		fmt.Println("Slices equal")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)

	slice2 := slice1[2:4]
	fmt.Println(slice2)
	fmt.Println("The capacity of slice2 is", cap(slice2))
	fmt.Println("The length of slice2 is", len(slice2))
}
