package main

/*
Arrays keep values of the SAME TYPE.
Arrays are FIXED SIZE.
*/

import (
	"fmt"
)

// Creates an array that can contain 3 items of the STRING type
// This array is empty when it's first created
// var names [3]string

// Create AND initialize an array of 3 strings
// states := [3]string{"California", "New York", "Colorado"}

// Update an item in an existing array
// states[2] = "Oregon"

var roles = [4]string{
	"guest",
	"user",
	"moderator",
	"admin",
}

func main() {
	fmt.Println(roles)
}
