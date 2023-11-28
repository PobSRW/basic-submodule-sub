package main

import "fmt"

func subProject() string {
	return "this is a sub project"
}

func main() {
	word := subProject()
	fmt.Println(word)

	// add comment to submodule from main module
	// add branch develop
}
