// variables
package main

import (
	"fmt"
)

func variables() {
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

	/*
		shortest
		in this case := is shorthand for var fun3 = false
		in general :=  no keyword var, no explicit type
		may appear only inside functions
		default type for built-in types is bool, rune, int, float64, complex128 or string
	*/
	fun3 := false
	fmt.Printf("Is Go fun? %t \n", fun3)
}

func multiplevariables() {
	/*
			Shortest
		   Define three integers without type "type" and without keyword "var", and initialize their values.
		   num1 is 1，num2 is 2，num3 is 3
	*/
	num1, num2, num3 := 1, 2, 3
	fmt.Printf("numbers are: %d %d %d \n", num1, num2, num3)

	/*
		Longer

	*/

	var real4, real5, real6 = 1.65, 3.14, 100.0
	fmt.Printf("real numbers are: %f %f %f \n", real4, real5, real6)
}

func blankVariable() {
	/*
		_ (blank) is a special variable name. Any value that is given to it will be ignored.
		Can be seen a lot about functions that return error.
		e.g result,_ := doSomethingThatCanReturnError()
	*/
	_, a := "unused", "used"
	fmt.Printf("variable a is %s \n", a)
}

const MAX_THREADS = 1000
const RED = "green"

//NOTE const FAIL := "fail" this won't compile, shorthand declaration is possible only in functions
