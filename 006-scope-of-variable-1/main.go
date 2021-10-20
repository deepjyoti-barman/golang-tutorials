package main

import "fmt"

func main() {
	// scope = where in your program you have access to any variable that you have declared, and what that value might be.
	// block level variables: They are only available inside the function they are declared
	var one = "One"
	fmt.Println(one)

	myFunc()
}

func myFunc() {
	var one = "1"
	fmt.Println(one)
}