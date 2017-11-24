// main
package main

import (
	"fmt"
)

func main() {

	//Extended Backus-Naur Form (EBNF)

	//longest variable declarartion
	var fun bool
	//variable assigment
	fun = true
	fmt.Printf("Is Go fun? %t \n", fun)

	//shorter
	// variable declaration + assigment
	var fun1 bool = false
	fmt.Printf("Is Go fun? %t \n", fun1)

	//shorter
	//type specification is not required
	var fun2 = true
	fmt.Printf("Is Go fun? %t \n", fun2)

	//shortest
	// := is shorthand for var fun3 = false
	//may appear only inside functions
	//default type for built-in types is bool, rune, int, float64, complex128 or string
	fun3 := false
	fmt.Printf("Is Go fun? %t \n", fun3)
}
