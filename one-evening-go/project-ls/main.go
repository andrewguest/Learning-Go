package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	files := listFiles("testdata")

	// Join each item of the Slice together, separated by a space
	// This converts the Slice into a string
	fmt.Print(strings.Join(files, " "))
}

func listFiles(dirname string) []string {
	var dirs []string

	files, _ := os.ReadDir(dirname)

	for _, f := range files {
		dirs = append(dirs, f.Name())
	}

	return dirs
}
