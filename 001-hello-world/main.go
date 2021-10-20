package main

import "fmt"

func main() {
	// package main - tells the Go compiler that the package should compile as an executable program instead of a shared library
	// i.e. executing the command 'go build <filename.go>' will produce a new executable file having the same name as the actual file that we have compiled
	// if you change the package name to something else it executing the command 'go build <filename.go>' won't generate an executable file and executing the command 'go run <filename.go>' will throw an error "package command-line-arguments is not a main package"
	fmt.Println("Hello world!")
}
