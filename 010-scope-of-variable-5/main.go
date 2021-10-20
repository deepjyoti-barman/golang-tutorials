package main

import (
	"fmt"
	"myapp/packageone"
)

func main() {
	firstString := packageone.PublicVar
	fmt.Println("From packageone, PublicVar =", firstString)

	// error: cannot refer to unexported name packageone.privateVar
	// secondString := packageone.privateVar
	// fmt.Println("From packageone, privateVar =", secondString)

	packageone.Exported()

	// error: cannot refer to unexported name packageone.notExported
	// packageone.notExported()
}