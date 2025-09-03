package main

import (
	"fmt"
	"os"
)

func main() {

	defer fmt.Println("Deferred statement")
	fmt.Println("Starting the main function")

	os.Exit(1)

	//fmt.Println("End of the main function")
}
