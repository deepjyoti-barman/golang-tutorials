package main

import "fmt"

// package level variable: package level variables are available to everywhere in the package where they are declared.
var one = "One"

func main() {
	// variable shadowing: block level variable takes precedence over package level variable (mostly as a good programming practice we should ignore this)
	var one = "This is a block level variable"
	fmt.Println(one)
	myFunc()
}

func myFunc() {
	fmt.Println(one)
}