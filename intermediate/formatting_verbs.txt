package main

import "fmt"

func main() {

	// --- General Formatting Verbs
	// %v	Prints the value in the default format
	// %#v	Prints the value in Go-syntax format (e.g. how it is printed in source code
	// %T	Prints the type of the value
	// %%	Prints the % sign

	//i := 1_5.5 // _ can be used to separate digits for readability
	//str := "Hello World!"
	//
	//fmt.Printf("%v\n", i)
	//fmt.Printf("%#v\n", i)
	//fmt.Printf("%T\n", i)
	//fmt.Printf("%v%%\n", i)
	//fmt.Printf("%v\n", str)
	//fmt.Printf("%#v\n", str)
	//fmt.Printf("%T\n", str)

	// --- 	Integer Formatting Verbs
	// %b	Base 2
	// %d	Base 10
	// %+d	Base 10 and always show sign
	// %o	Base 8
	// %O	Base 8, with leading 0o
	// %x	Base 16, lowercase
	// %X	Base 16, uppercase
	// %#x	Base 16, with leading 0x
	// %4d	Pad with spaces (width 4, right justified)
	// %-4d Pad with spaces (width 4, left justified)
	// %04d Pad with zeroes (width 4)

	//i2 := 255
	//fmt.Printf("%b\n", i2)
	//fmt.Printf("%d\n", i2)
	//fmt.Printf("%+d\n", i2)
	//fmt.Printf("%o\n", i2)
	//fmt.Printf("%O\n", i2)
	//fmt.Printf("%x\n", i2)
	//fmt.Printf("%X\n", i2)
	//fmt.Printf("%#x\n", i2)
	//fmt.Printf("%#X\n", i2)
	//fmt.Printf("%4d\n", i2)
	//fmt.Printf("%-4d\n", i2)
	//fmt.Printf("%02d\n", i2)

	// --- String Formatting Verbs
	// %s	Prints the value as a plain string
	// %q	Prints the value as a double-quoted string
	// %8s	Prints the value as a plain string (width 8, right justified)
	// %-8s	Prints the value as a plain string (width 8, left justified)
	// %x	Prints the value as a hex dump of byte values
	// % x	Prints the value as a hex dump with spaces

	//text := "World"
	//fmt.Printf("%s\n", text)
	//fmt.Printf("%q\n", text)
	//fmt.Printf("%8s\n", text)
	//fmt.Printf("%-8s\n", text)
	//fmt.Printf("%x\n", text)
	//fmt.Printf("% x\n", text)

	// --- Boolean Formatting Verbs
	// %t	Value of the boolean operator in true or false format (same as using %v)
	//
	//t := true
	//f := false
	//fmt.Printf("%t\n", t)
	//fmt.Printf("%v\n", t)
	//fmt.Printf("%t\n", f)

	// --- Float Formatting Verbs
	// %e	Scientific notation with 'e' as exponent
	// %f	Decimal point, no exponent
	// %.2f	Default width, precision 2
	// %6.2f	Width 6, precision 2
	// %g	Exponent as needed, only necessary digits

	f := 918.026
	fmt.Printf("%e\n", f)
	fmt.Printf("%f\n", f)
	fmt.Printf("%.2f\n", f)
	fmt.Printf("%10.2f\n", f)
	fmt.Printf("%g\n", f)
}
