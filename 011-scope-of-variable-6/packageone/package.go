package packageone

import "fmt"

var PackageVar = "PACKAGE_VAR"

func PrintMe(myVar, blockVar string) {
	fmt.Println("myVar =", myVar)
	fmt.Println("blockVar =", blockVar)
	fmt.Println("PackageVar =", PackageVar)
}