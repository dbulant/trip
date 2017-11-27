// main
package main

import . "fmt"

func main() {
	Printf("!!Mention go playground!! \n")
	//Printf("Any variable that begins with a capital letter means it will be visible to other packages, otherwise. \n")
	//Printf("The same rule applies for functions and constants, no public or private keyword exists in Go. \n")

	cache := NewPersonsCache()
	server := NewRestServer(cache)
	server.Listen(":9876")
}
