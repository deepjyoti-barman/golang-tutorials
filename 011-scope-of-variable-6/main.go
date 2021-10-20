package main

import (
	"myapp/packageone"
)

// package level variable
var myVar = "PACKAGE_VAR"

func main() {
	var blockVar = "BLOCK_VAR"
	packageone.PrintMe(myVar, blockVar)
}