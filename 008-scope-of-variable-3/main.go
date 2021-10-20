package main

import "fmt"

// package level variable: package level variables are available to everywhere in the package where they are declared.
var one = "One"

func main() {
	fmt.Println(one)
	myFunc()
}

func myFunc() {
	fmt.Println(one)
}