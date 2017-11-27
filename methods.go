// methods
package main

import (
	"fmt"
	"math"
)

/*
"A method is a function with an implicit first argument, called a receiver."
Syntax of method
func (r ReceiverType) funcName(parameters) (results)
receiver = "this" or "self"
*/

type Circle struct {
	radius float64
}

// method
func (this *Circle) Area() float64 {
	return this.radius * this.radius * math.Pi
}

type Rectangle struct {
	width, height float64
}

// method
func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

/*
If the name of methods are the same but they don't share the same receivers, they are not the same.
Methods are able to access fields within receivers.
Use . to call a method in the struct, the same way fields are called.
*/
func areas() {
	c1 := Circle{10}
	c2 := Circle{25}
	r1 := Rectangle{9, 4}
	r2 := Rectangle{12, 2}

	fmt.Println("Area of c1 is: ", c1.Area())
	fmt.Println("Area of c2 is: ", c2.Area())
	fmt.Println("Area of r1 is: ", r1.Area())
	fmt.Println("Area of r2 is: ", r2.Area())
}

//see  methods in structures.go
func overriding() {
	sam := Employee{Human: Human{name: "Sam", age: 45, phone: "111-888-XXXX"}, company: "factory", phone: "X88-166"}
	mark := Student{Human: Human{name: "Mark", age: 25, weight: 120}, specialty: "Chemistry"}

	sam.SayHello()
	mark.SayHello()
}
