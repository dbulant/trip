// functions
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
Generic declaration
func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
    // function body
    // multi-value return
    return value1, value2
}

NOTE
-Use keyword func to define a function funcName.
-Functions have zero, one or more than one arguments.
The argument type comes after the argument name and arguments are separated by ,.
-Functions can return multiple values.
-You can omit their names and use their type only.
-If there is only one return value and you omitted the name, you don't need brackets for the return values.
-If the function doesn't have return values, you can omit the return parameters altogether.
-If the function has return values, you have to use the return statement somewhere in the body of the function.
*/

func SumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}

func SumAndProductNamed(a int, b int) (sum int, product int) {
	sum = a + b
	product = a * b
	return
}

func variadicFunction(ints ...int) {
	for _, i := range ints {
		fmt.Println(i)
	}
}

//NOTE Go passes everything by value
func addBad(a int) {
	a = a + 1
	return
}

//Need to pass by pointer here
func addGood(a *int) {
	*a = *a + 1
}

//Typical swap example
func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func addition() {
	a := 10
	fmt.Println("a is", a)

	addBad(a)
	fmt.Println("a after addBad(a) is", a)

	addGood(&a)
	fmt.Println("a after addGood(&a) is", a)
}

func swapping() {
	a, b := 10, 5
	fmt.Printf("a is: %d, b is: %d \n", a, b)
	swap(&a, &b)
	fmt.Printf("a is: %d, b is: %d \n", a, b)
}

func deferFunction() {
	//Close the file when you’re done (usually this would be scheduled immediately after Opening with defer).
	//lots of err variables here -> it means after very error check f.Close() should be called
	//using defer we schedule it to the end of function

	f, err := os.Open("/tmp/dat")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	o2, err := f.Seek(6, 0)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	o3, err := f.Seek(6, 0)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	fmt.Printf("5 bytes: %s\n", string(b4))
}

//You can have more defers in function. They execute in reverse order.
func moreDefersFunction() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("called defer with value %d \n ", i)
	}
}

//Functions are also variables in Go, we can use type to define them.
//Functions that have the same signature can be seen as the same type.
//type typeName func(input1 inputType1 , input2 inputType2 [, ...]) (result1 resultType1 [, ...])
//as a result of this feature,we can pass functions as values (not in C style)
type testInteger func(int) bool

func isOdd(integer int) bool {
	return integer%2 != 0
}

func isEven(integer int) bool {
	return integer%2 == 0
}

//These two ways of passing function to function are equal.
//func filter(slice []int, f testInteger) []int OR func filter(slice []int, f func(int) bool) []int
func filter(slice []int, f testInteger) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func passingFunctionAsValue() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	odd := filter(s, isOdd)
	even := filter(s, isEven)

	fmt.Printf("slice is %+v \n", s)
	fmt.Println("Odd elements of slice are: ", odd)
	fmt.Println("Even elements of slice are: ", even)
}

//Panic is a built-in function to break the normal flow of programs and get into panic status.
//In this status, only defer functions will continue to execute.
//Recover is a built-in function to recover goroutines from panic status
func panicFunction() {
	ev := os.Getenv("NOTPOSSIBLE")
	if ev == "" {
		panic("no value for $NOTPOSSIBLE")
	}
}

/*Go has two special functions which are called main and init, where init can be used in all packages and main can only be used in the main package.
These two functions are not able to have arguments or return values.
Even though we can write many init functions in one package, it is strongly recommended to write only one init function for each package.
Go programs will call init() and main() automatically, so you don't need to call them by yourself.
For every package, the init function is optional, but package main has one and only one main function.+
init() is executed after importing package
*/
