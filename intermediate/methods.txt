package main

import "fmt"

type Shape struct {
	Rectangle
}

type Rectangle struct {
	length float64
	width  float64
}

// Area Method with value receiver. Used when we don't modify the underlying data
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Scale Method with pointer receiver. Will modify the underlying data
func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

func main() {
	rectangle := Rectangle{length: 10, width: 9}
	area := rectangle.Area()
	fmt.Println("Area is", area)
	rectangle.Scale(2.0)
	fmt.Println("Area is", rectangle.Area())

	num := MyInt(-5)
	num1 := MyInt(10)
	fmt.Println(num.IsPositive())
	fmt.Println(num1.IsPositive())

	fmt.Println(num.welcomeMessage())

	s := Shape{
		Rectangle: Rectangle{length: 10, width: 9},
	}
	fmt.Println(s.Area()) // Method is promoted to the struct from the Anonymous type
}

type MyInt int

// IsPositive Method on a user defined type. Not just used on structs
func (m MyInt) IsPositive() bool {
	return m > 0
}

// Don't need to use an instance.
func (MyInt) welcomeMessage() string {
	return "Welcome to MyInt Type"
}
