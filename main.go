package main

import (
	"fmt"
	t "time"
)

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

	/* variable named value,

	var radius = 6
	radius := 6

	*/

	// Basic Data Types - int, float, complex.

	/*
		int8, uint8, int16, uint16, int32 , uint32, int64, uint64, int, uint, uintptr
		float32, float64
		complex64, complex128
	*/

	var publisher, writer, artist, title string
	var year uint
	var pageNumber uint
	var grade float32

	title = "Mr.GoToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	publisher = "Dizzybooks Publishing Inc"
	year = 1997
	pageNumber = 14
	grade = 6.5

	fmt.Println(title, "written by", writer, "drawn by", artist)
	fmt.Println("Publisher:", publisher, "Pgs:", pageNumber, "Yr:", year, "Grade:", grade)
	fmt.Println()

	title = "Epic Vol. 1"
	writer = "Ryan N. Shawn"
	artist = "Phoebe Paperclips"
	publisher = "Dizzybooks Publishing Inc"
	year = 2013
	pageNumber = 160
	grade = 9.0

	fmt.Println(title, "written by", writer, "drawn by", artist)
	fmt.Println("Publisher:", publisher, "Pgs:", pageNumber, "Yr:", year, "Grade:", grade)
	fmt.Println()

	// Println VS Print

	fmt.Println("Let's first see how", "the Println() method works.")
	fmt.Println("Notice that each statement adds a newline for us.")
	fmt.Println("There's also a default space", "between the string arguments.")
	fmt.Print("Print", "is", "different") // no space
	fmt.Print("See?")                     // adds to last possible place

	// %v - Interpolation/Verbs

	specialNum := 42
	fmt.Printf("This value's type is %T.", specialNum) // Prints value type 'int'

	quote := "To do or not to do"
	fmt.Printf("This value's type is %T.", quote) // Prints value type 'string'

	animal1 := "cat"
	animal2 := "dog"
	fmt.Printf("Are you a %v or a %v person?", animal1, animal2) // v = 'verb'

	Num1 := 18
	Num2 := 3.8
	fmt.Printf("You must be %d years old to vote.", Num1) // d = 'int to string'
	fmt.Printf("You're averaging: %f.", Num2)             // f = 'int to string (float)'
	fmt.Printf("You're averaging: %.2f.", Num2)           // f = 'int to string (precision float)'

	fmt.Println("How are you doing?")
	var response string
	fmt.Scan(&response)
	fmt.Printf("I'm %v.", response)

	heistReady := false
	if heistReady {
		fmt.Println("Let's go!")
	} else if heistReady == true {
		fmt.Println("else if is working!")
	} else {
		fmt.Println("Act normal")
	}

	// '!' reverses the logic

	heistReady1 := true // Var set to true, returns opp, logic prints false
	if !heistReady1 {
		fmt.Println("Let's go!")
	} else {
		fmt.Println("Act normal")
	}

	// switch in place of if/else if/ else
	someVariable := 2
	switch someVariable {
	case 1:
		fmt.Println("I got first place!")
	case 2:
		fmt.Println("I got second place!")
	case 3:
		fmt.Println("I got third place!")
	}

}
