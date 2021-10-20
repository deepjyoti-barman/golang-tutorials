package packageone

import "fmt"

// a private variable starts with a lowercase character and is only available inside this package.
var privateVar = "I am private"

// a public variable starts with a uppercase character, it's exported - it's available outside of this package (to any other package that imports this package)
var PublicVar  = "I am public (or exported)"

// a private function starts with a lowercase character and is only available inside this package
func notExported() {
	fmt.Println("From notExported()")
}

// a public function starts with a uppercase character, it's exported - it's available outside of this package (to any other package that imports this package)
func Exported() {
	fmt.Println("From Exported()")
}