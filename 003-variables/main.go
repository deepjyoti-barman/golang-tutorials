package main

import "fmt"

func main() {
	/*
		rules of declaring a variable:
			- variable name can't start with a number (e.g. 1number)
			- variable name can either start with a character (uppercase or lowercase) or an underscore (e.g. _firstNumber)
	*/

	// first way - declare, then assign (two steps)
	var firstNumber int
	firstNumber = 2

	// second way - declare type and name and assign value
	var secondNumber = 5

	// third way - one step variable: declare name, assign value, and let Go compiler figure out the type
	thirdNumber := 7

	// The default variable specifier is '%v', which is a catch-all for most variable types.
	fmt.Printf("firstNumber = %v, secondNumber = %v and thirdNumber = %v\n", firstNumber, secondNumber, thirdNumber)

	// Printing the default values
	var strVar string
	var intVar int
	var floatVar float64
	var complexVar complex64
	var booleanVar bool

	fmt.Println("\nPrinting the default values of all data-types:")
	fmt.Println("strVar =", strVar)
	fmt.Println("intVar =", intVar)
	fmt.Println("floatVar =", floatVar)
	fmt.Println("complexVar =", complexVar)
	fmt.Println("booleanVar =", booleanVar)
}
