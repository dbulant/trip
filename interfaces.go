// interfaces
package main

import "fmt"

type HumanBeing struct {
	name  string
	age   int
	phone string
}

type UniversityStudent struct {
	HumanBeing
	school string
	loan   float32
}

type Employe struct {
	HumanBeing
	company string
	money   float32
}

//an interface is a set of methods that we use to define a set of actions.
// define interfaces
type Men interface {
	SayHi()
	Sing(lyrics string)
}

type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}

func (p HumanBeing) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", p.name, p.phone)
}

func (p HumanBeing) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

func (p HumanBeing) TakeNap(amount int) {
	fmt.Printf("Sleeping for %d ", amount)
}

// Employe overloads SayHi
func (e Employe) SayHi() {
	fmt.Printf("Hi, I am employee %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}

func (e Employe) SpendSalary(amount float32) {
	e.money -= amount
}

func (e Employe) Sing(song string) {
	fmt.Printf("not in mood to sing %s \n", song)
}

func (s UniversityStudent) SayHi() {
	fmt.Printf("Hi, I am university student called %s and my loan is %g \n", s.name, s.loan)
}

func (s UniversityStudent) Sing(song string) {
	fmt.Printf("trala la la ala %s \n", song)
}

func (s UniversityStudent) BorrowMoney(amount float32) {
	s.loan += amount // (again and again and...)
}

func interfaces() {
	mike := UniversityStudent{HumanBeing{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := UniversityStudent{HumanBeing{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employe{HumanBeing{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := Employe{HumanBeing{"Sam", 36, "444-222-XXX"}, "Things Ltd.", 5000}

	// define interface i
	var i Men
	//WHY IT HAS TO BE USED ON NON-POINTER METHODS?,
	//i can store Student
	i = mike
	i.SayHi()
	i.Sing("November rain")

	//i can store Employee
	i = tom
	i.SayHi()
	i.Sing("Born to be wild")

	// slice of Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	// these three elements are different types but they all implemented interface Men
	//overriding virtual methods (C++ style)
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
	}
}
