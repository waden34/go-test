package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	message := "Hello, Go!"
	message1 := "Hello\tGo!"
	rawMessage := `Hello\nGo` // ` is string literal. No escape sequences

	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(rawMessage)

	fmt.Println("Length of message: ", len(message))

	fmt.Println("First character in message var:", message[0]) // ASCII

	greeting := "Hello"
	name := "Jeremiah"
	fmt.Println(greeting + name)

	str1 := "Apple"
	str := "apple"
	str2 := "banana"
	str3 := "app"
	fmt.Println(str1 < str2) // < is comparing ascii value of runes
	fmt.Println(str3 < str1)
	fmt.Println(str > str1)
	fmt.Println(str > str3)

	for _, c := range message {
		//fmt.Printf("Character at index %d is %c\n", i, c)
		fmt.Printf("%v\n", c) // %v is ascii value of the rune. %x is hex value, %c is string value
	}

	fmt.Println("Rune count in greeting:", utf8.RuneCountInString(greeting))

	greetingWithName := greeting + " " + name
	fmt.Println(greetingWithName)

	var ch rune = 'a' // ' is rune
	jch := '日'
	fmt.Println(ch)
	fmt.Println(jch)

	fmt.Printf("%c\n", ch)
	fmt.Printf("%c\n", jch)

	cstr := string(ch)
	fmt.Println(cstr)
	fmt.Printf("Type of cstr is %T\n", cstr)

	const NIHONGO = "日本語"
	fmt.Println(NIHONGO)

	hello := "こんにちは"
	for _, runeValue := range hello {
		fmt.Printf("%c\n", runeValue)
	}

	r := '😁'
	fmt.Printf("%v\n", r)
	fmt.Printf("%c\n", r)
}
