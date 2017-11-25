// flowcontrol
package main

import (
	. "fmt"
	"math/rand"
	"time"
)

func generateRandomNumber() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(10)
}

func ifcondition() {
	Printf("No parentheses required \n")
	rn := generateRandomNumber()
	if rn < 5 {
		Printf("generated number is less than 5, %d \n", rn)
	} else {
		Printf("generated number is more than 5, %d \n", rn)
	}

	Printf(" Go can have one initialization statement before the conditional statement. \n")
	if g := generateRandomNumber(); g < 5 {
		Printf("generated number is less than 5, %d \n", g)
	} else if g == 5 {
		Printf("generated number is 5, %d \n", g)
	} else {
		Printf("generated number is more than 5, %d \n", g)
	}
	//NOTE visibility of g is only in if-else block, using g outside causes compilation error
}

func gotocondition() {
	i := 0
Here: // label ends with ":"
	Println(i)
	i++
	if i == 10 {
		return
	}
	goto Here // jump to label "Here"
}

func forcycle() {
	//for expression1; expression2; expression3

}

func breakforcycle() {
	for i := 10; i > 0; i-- {
		if i == 5 {
			break // or continue
		}
		Println(i)
	}
}

func switchcondition() {
	i := generateRandomNumber()
	//NOTE no break after case is required
	switch i {
	case 1:
		Println("i is equal to 1")
	case 2, 3, 4: //more values on same line
		Println("i is equal to 2, 3 or 4")
	case 10:
		Println("i is equal to 10")
	default:
		Println("All I know is that i is an integer")
	}
}

func switchfalthrough() {
	integer := generateRandomNumber()
	//NOTE use fallthrough to continue matching more cases
	switch integer {

	case 1:
		Println("integer <= 1")
		fallthrough
	case 2:
		Println("integer <= 2")
		fallthrough
	case 3:
		Println("integer <= 3")
		fallthrough
	case 4:
		Println("integer <= 4")
		fallthrough
	case 5:
		Println("integer <= 5")
		fallthrough
	case 6:
		Println("integer <= 6")
		fallthrough
	case 7:
		Println("integer <= 7")
		fallthrough
	case 8:
		Println("integer <= 8")
		fallthrough
	default:
		Println("default case")
	}
}

func continueforcycle() {
	for i := 10; i > 0; i-- {
		if i == 5 {
			continue // or continue
		}
		Println(i)
	}
}

func sum1(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}

func sum2(n int) int {
	//Go has no while
	//it's behavior using for
	sum := 1
	for sum < n {
		sum += sum
	}
	return sum
}
