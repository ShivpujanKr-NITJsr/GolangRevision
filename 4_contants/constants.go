package main

import "fmt"

const age = 30  // constant integer
const p int = 8 // constant integer with type

// name:="golang"   // cant declare out of function
func main() {
	const name string = "golang" // constant string
	// name = "golang" // cannot assign a new value to a constant

	// constant grouping

	const (
		port = 5000
		host = "localhost"
	)    // later cant reassigned

	fmt.Println(port, host)

}