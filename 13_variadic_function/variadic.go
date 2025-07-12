package main

import "fmt"

// Variadic functions allow you to pass a variable  number of arguments to a function.

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
// interface{} can be used to accept any type of argument, but it is not type-safe.
// The `...` syntax in the function definition indicates that it can accept a variable number of arguments.

func main() {
	// fmt.Println(1, 2, 3, 5, "hello")
	// result := sum(3,4,5,6,7,8,9,10) // passing variable number of arguments
	// fmt.Println("Sum:", result) // prints Sum: 52

	num := []int{1, 2, 3, 4, 5}
	result := sum(num...) // using the spread operator to pass the slice as individual arguments
	fmt.Println("Sum:", result) // prints Sum: 15

	
}
