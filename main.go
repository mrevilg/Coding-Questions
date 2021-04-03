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

	fmt.Println("Using the 'go doc' command is helpful")
	fmt.Println("You can find out more about a package")
	fmt.Println("Or about a function inside the package")
	fmt.Println("Try it out in the terminal!")
	fmt.Println("go doc fmt.Println") // or time.Now

}