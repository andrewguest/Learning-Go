package main

/*
Slices are "dynamic arrays".
You don't declare how many elements they have, because they can expand and shrink.
Slices can still only contain one data type and that type has to be declared.
*/

import "fmt"

var actions = []string{"create", "read", "update", "delete"}

func main() {
	fmt.Println(actions)
}
