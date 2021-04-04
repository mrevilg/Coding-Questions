package main

import "fmt"
import t "time"

// This is a Comment

/* 
This is a comment block
*/

func main() {
	fmt.Println("Hello World")
	fmt.Println(t.Now())
	fmt.Println(2235 * 1231)
	fmt.Println("Using the 'go doc' command is helpful")
	fmt.Println("You can find out more about a package")
	fmt.Println("Or about a function inside the package")
	fmt.Println("Try it out in the terminal!")
	fmt.Println("go doc fmt.Println") // or time.Now

	// cannot be update whilst running
	const constant = "This is a const being called" 
	fmt.Println(constant)

	// literal unnamed value
	fmt.Println("PI = ", 3.14159)
 
	// constant named value
	const pi = 3.14159
	
	// variable named value
	var radius = 6

	// Basic Data Types - int, float, complex
	/*
	int8, uint8, int16, uint16, int32 , uint32, int64, uint64, int, uint, uintptr
	float32, float64
	complex64, complex128
	*/
}