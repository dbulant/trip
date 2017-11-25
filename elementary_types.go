// elementary_types
package main

import (
	"errors"
	"fmt"
)

func booleans() {
	fmt.Printf("Booleans function \n")
	/*
		You cannot convert variables' type between number and boolean!
		NOTE below throws compile error
		t := true
		var i int = int(t)
	*/
}

func integerTypes() {
	fmt.Printf("Cannot assign values between types, explicit casting is required!! \n")
	var r rune = 97
	var i8 int8 = (1 << 7) - 1
	var i16 int16 = (1 << 15) - 1
	var i32 int32 = (1 << 31) - 1
	var i64 int64 = (1 << 63) - 1
	var b byte = 1<<8 - 1
	var ui8 uint8 = 1<<8 - 1
	var ui16 uint16 = 1<<16 - 1
	var ui32 uint32 = 1<<32 - 1
	var ui64 uint64 = 1<<64 - 1
	fmt.Printf("rune is just alias for int32 %c \n", r)
	fmt.Printf("byte is just alias for uint8 %d \n", b)
	fmt.Printf("Max positive values for signed types are: %d  %d  %d  %d \n", i8, i16, i32, i64)
	fmt.Printf("Max negative values for signed types are: %d  %d  %d  %d \n", -i8-1, -i16-1, -i32-1, -i64-1)
	fmt.Printf("Max values for unsigned types are: %d  %d  %d  %d \n", ui8, ui16, ui32, ui64)
	//NOTE this causes compile errors
	//i8 = ui8
	//You have to cast it explicitly.
	ui64 = uint64(ui32)
}

func floatingPointAndComplexTypes() {
	var f1 float32 = 1<<32 - 1
	var f2 float64 = 1<<500 - 1

	var c1 complex64 = 24i
	var c2 complex128 = -5 + 12i
	fmt.Printf("There is no default float type, only float32 and float64!! \n")
	fmt.Printf("f1 is %g f2 is %g \n", f1, f2)
	fmt.Printf("c1 is %g c2 is %g \n", c1, c2)
}

func stringType() {
	var emptyString string                 // basic form to define string
	var hello string = "hi"                // define a string with empty string
	no, yes, maybe := "no", "yes", "อาจจะ" // brief statement
	russianHello := "здравствуй"
	hello = "Bonjour" // basic form of assign values
	fmt.Printf("UTF8 default support %s %s %s %s %s \n", no, yes, maybe, russianHello, hello)
	fmt.Printf("empty string has value %s \n", emptyString)

	/*
		It's impossible to change string values by index.
		NOTE compilation error
		changeMe := "changeMe"
		changeMe[0] = 'C'
	*/

	//If you really want to change string value by index.
	s := "hello"
	c := []byte(s) // convert string to []byte type
	c[0] = 'c'
	s2 := string(c) // convert back to string type
	fmt.Printf("%s\n", s2)

	s = "hello "
	m := " world"
	a := s + m
	fmt.Printf("%s\n", a)

	// you cannot change string values by index, but you can get values instead.
	a = s + m[:]
	fmt.Printf("%s\n", a)
}

func errorType() {
	err := errors.New("WTF! What a Terrible Failure! \n")
	if err != nil {
		fmt.Print(err)
	}
}

func iotaEnumerate() {
	const (
		x = iota // x == 0
		y = iota // y == 1
		z = iota // z == 2
		w        // If there is no expression after the constants name, it uses the last expression,
		//so it's saying w = iota implicitly. Therefore w == 3, and y and z both can omit "= iota" as well.
	)

	const v = iota // once iota meets keyword `const`, it resets to `0`, so v = 0.

	const (
		e, f, g = iota, iota, iota // e=0,f=0,g=0 values of iota are same in one line.
	)
	fmt.Printf("x: %d y: %d z: %d w: %d \n", x, y, z, w)
	fmt.Printf("v: %d \n", v)
	fmt.Printf("e: %d f: %d g: %d \n", e, f, g)
}

func arrayType() {
	//var arr [n]type n is number of elements
	//array of 3 ints
	var ai [3]int
	ai[0] = 9
	ai[1] = 3
	ai[2] = 6
	fmt.Printf("array is %v \n", ai)

	//using :=
	ai1 := [3]int{1, 2, 3}
	fmt.Printf("array is %v \n", ai1)

	//only first three elements are assigned, rest has default value of 0 (default value of its type)
	ai2 := [10]int{4, 5, 6}
	fmt.Printf("array is %v \n", ai2)

	// use `…` (ellipsis) to replace the length parameter and Go will calculate it for you.
	ai3 := [...]int{4, 5, 6}
	fmt.Printf("array is %v \n", ai3)
}

func sliceType() {
	/*
		https://blog.golang.org/go-slices-usage-and-internals
			slice "==" "dynamic array"
			Slice is descriptor of an array segment.
			It contains pointer to an underlying array

			length - The length is the number of elements referred to by the slice.
			capacity - The capacity is the number of elements in the underlying array (beginning at the element referred to by the slice pointer)

			Slice is declared same way as array, except element count.
			Slice is created with built-in function make
			func make([]T, len, cap) []T
	*/

	a := [3]int{1, 2, 3}
	fmt.Printf("This is array %T \n", a)

	ns := []int{1, 2, 3}
	fmt.Printf("This is slice %T \n", ns)

	//Default values are used, i.e for int it is 0
	s := make([]int, 5, 5)
	fmt.Printf("This is slice %T %v \n", s, s)
	fmt.Printf("length is %d, capacity is %d \n", len(s), cap(s))
	//iterating over slice usibg range keyword
	for i, _ := range s {
		s[i] = i + 1
	}
	fmt.Printf("Slice is %v \n", s)
	//2,3 include, 4 excluded, half open range
	r := s[2:4]
	fmt.Printf("Slice r is %v \n", r)

	//Slicing does not copy the slice's data.
	//It creates a new slice value that points to the original array.
	//This makes slice operations as efficient as manipulating array indices.
	//Therefore, modifying the elements (not the slice itself) of a re-slice modifies the elements of the original slice.
	r[0] = 1
	fmt.Printf("Slice s is %v \n", s)
	fmt.Printf("Slice r is %v \n", r)

	//The append function appends the elements x to the end of the slice s, and grows the slice if a greater capacity is needed.
	//func append(s []T, x ...T) []T
	sr := make([]int, 0)
	sr = append(sr, 1001)
	//If we want to append another slice, ... (ellipsis) is required
	sr = append(sr, s...)
	sr = append(sr, r...)
	fmt.Printf("Slice sr is %v \n", sr)

}

func mapType() {
	//map behaves like a dictionary in Python. Form map[keyType]valueType is used to define it.
	// use string as the key type, int as the value type, and `make` initialize it.
	var emptyMap map[string]int
	emptyMap = make(map[string]int)
	fmt.Printf("emptyMap is %+v \n", emptyMap)
	// another way to define map
	numbers := make(map[string]int)
	numbers["one"] = 1 // assign value by key
	numbers["ten"] = 10
	numbers["three"] = 3

	//for range to iterate over map
	for k, v := range numbers {
		fmt.Println("numbers's key:", k)
		fmt.Println("numbers's val:", v)
	}

	//if you want only value
	for _, v := range numbers {
		fmt.Println("numbers's val:", v)
	}

	//Get value from map
	fmt.Printf("The third number is: %d \n", numbers["three"])

	/*
		Map is disorderly. Everytime you print map you will get different order. It's impossible to get values by index -you have to use key.
		Map doesn't have a fixed length. It's a reference type just like slice.
		len works for map also. It returns how many keys that map has.
		It's quite easy to change the value through map. e.g Simply use numbers["one"]=11 to change the value of key one to 11
	*/
	numbers["eleven"] = 11
	numbers["three"] = 3
	fmt.Printf("map is %+v \n", numbers)

	/*
		You can use form key:val to initialize map's values, and map has built-in methods to check if the key exists.
		Use delete to delete an element in map.
	*/

	// Initialize a map
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	// map has two return values. For the second return value, if the key doesn't
	//exist，'ok' returns false. It returns true otherwise.
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}

	delete(rating, "C") // delete element with key "c"

	/*
		map is a reference type. If two maps point to same underlying data, any change will affect both of them.
	*/
	m := make(map[string]string)
	m["Hello"] = "Bonjour"
	m1 := m
	m1["Hello"] = "Salut" // now the value of m["hello"] is Salut
	fmt.Printf("map m is %+v \n", m)
	fmt.Printf("map m1 is %+v \n", m1)
	m["Hello"] = "สวัสดี"
	fmt.Printf("map m is %+v \n", m)
	fmt.Printf("map m1 is %+v \n", m1)
}
