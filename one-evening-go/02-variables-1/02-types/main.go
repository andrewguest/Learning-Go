package main

/*
Data types:
	int = An integer number. 0, 42, -200
	bool = true or false
	float64 = A decimal number. 1.23, 3.1145
	string = Text. "Alice", "Bob is nice"
		Strings are defined using double quotes, not single quotes.
		Single quotes are used to define a BYTE or RUNE
*/

import "fmt"

func main() {
	// Declare
	var daysInDecember int
	// Assign
	daysInDecember = 31

	// Declare and assign
	var pi = 3.1415

	// Also declare and assign
	learningGo := true

	fmt.Println("days in december:", daysInDecember)
	fmt.Println("pi:", pi)
	fmt.Println("learning go:", learningGo)
}
