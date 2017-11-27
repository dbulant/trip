// structures
package main

import "fmt"

//typical definition
type person struct {
	name string
	age  int
}

func basicStructUsage() {
	var P person // p is person type

	P.name = "Yen"                                  // assign "Yen" to the field 'name' of p
	P.age = 28                                      // assign 25 to field 'age' of p
	fmt.Printf("The person's name is %s\n", P.name) // access field 'name' of p
}

func structInitialization() {
	//initialize by default order of fields
	P := person{"Tom", 25}
	//Use the format field:value to initialize the struct without order
	Q := person{age: 24, name: "Bob"}
	//Define an anonymous struct, then initialize it
	R := struct {
		name string
		age  int
	}{"Amy", 18}
	fmt.Printf("%+v %+v %+v \n", P, Q, R)
}

//Embedded fields in struct
type Human struct {
	name   string
	age    int
	weight int
	phone  string
}

func embeddedFields1() {
	// instantiate and initialize a student
	mark := Student{Human: Human{name: "Mark", age: 20, weight: 120}, specialty: "Biology"}
	//mark := Student{Human{"Mark", 25, 120}, nil, 0, "Biology"} <- this works

	// Print structure
	fmt.Printf("%+v \n", mark)

	// modify mark's specialty
	mark.specialty = "chemistry"
	fmt.Println("Mark changed his specialty")
	fmt.Println("His specialty is ", mark.specialty)

	fmt.Println("Mark become old. He is not an athlete anymore")
	mark.age = 46
	mark.weight += 60
	fmt.Println("His age is", mark.age)
	fmt.Println("His weight is", mark.weight)

	fmt.Println("Mark become too old. We have to use embedded fields access.")
	mark.age = 60
	mark.weight += 80
	fmt.Println("His age is", mark.Human.age)
	fmt.Println("His weight is", mark.Human.weight)
}

type Skills []string

type Student struct {
	Human     // struct as embedded field
	Skills    // string slice as embedded field
	int       // built-in type as embedded field
	specialty string
}

func (s *Student) SayHello() {
	fmt.Printf("Hello, I am %s my specialization is: %s \n", s.name, s.specialty)
}

func embeddedFields2() {
	// initialize Student Jane
	jane := Student{Human: Human{"Jane", 35, 100, "333-555-666"}, specialty: "Biology"}
	// access fields
	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println("Her weight is ", jane.weight)
	fmt.Println("Her specialty is ", jane.specialty)

	// modify value of skill field
	jane.Skills = []string{"anatomy"}
	fmt.Println("Her skills are ", jane.Skills)
	fmt.Println("She acquired two new ones ")
	jane.Skills = append(jane.Skills, "physics", "chemistry")
	fmt.Println("Her skills now are ", jane.Skills)
	// modify embedded field
	jane.int = 3
	fmt.Println("Her preferred number is ", jane.int)
}

type Employee struct {
	Human
	company string
	phone   string // phone in employee
}

func (e *Employee) SayHello() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //You can split into 2 lines here.
}

func outerInnerAccess() {
	Bob := Employee{Human{"Bob", 34, 200, "777-444-XXXX"}, "Designer", "333-222"}

	fmt.Println("Bob's work phone is:", Bob.phone)
	fmt.Println("Bob's personal phone is:", Bob.Human.phone)
}
