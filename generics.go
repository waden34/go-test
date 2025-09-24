package main

import "fmt"

func swap[T any](a, b T) (T, T) {
	return b, a
}

type Stack[T any] struct { // Stack is like a LIFO queue
	elements []T
}

func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1] // nothing before the : means we start at the start, end at the last one
	return element, true
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

func (s Stack[T]) printAll() {
	if len(s.elements) == 0 {
		fmt.Println("The stack is empty")
		return
	}
	fmt.Println("The stack elements")
	for _, element := range s.elements {
		fmt.Println("\t", element)
	}
	fmt.Println()
}

func main() {
	//x, y := 1, 2
	//x, y = swap(x, y)
	//fmt.Println(x, y)
	//
	//x1, y1 := "John", "Mary"
	//x1, y1 = swap(x1, y1)
	//fmt.Println(x1, y1)

	intStack := Stack[int]{}
	intStack.push(1)
	intStack.push(2)
	intStack.push(3)
	fmt.Println(intStack.pop())
	intStack.printAll()
	fmt.Println(intStack.pop())
	fmt.Println("Is stack empty", intStack.isEmpty())
	fmt.Println(intStack.pop())
	fmt.Println("Is stack empty", intStack.isEmpty())

	stringStack := Stack[string]{}
	stringStack.push("hello")
	stringStack.push("world")
	stringStack.push("john")
	stringStack.printAll()
	fmt.Println(stringStack.pop())
	fmt.Println("Is stack empty", stringStack.isEmpty())
	fmt.Println(stringStack.pop())
	fmt.Println("Is stack empty", stringStack.isEmpty())
	fmt.Println(stringStack.pop())
	fmt.Println("Is stack empty", stringStack.isEmpty())
}
