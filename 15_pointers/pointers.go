package main

import "fmt"

// by value pass
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("Inside changeNum:", num)
// }

// by reference pass
func changeNum(num *int) {
	*num = 5 // dereferencing the pointer to change the value at that address
	fmt.Println("Inside changeNum:", *num) // prints 5
}

func main() {

	// num := 1
	// changeNum(num)                       // passing by value, num is copied

	// fmt.Println("Memory address of num:", &num) // prints memory address of num
	// fmt.Println("After changeNum:", num) // prints 1, original num is unchanged
	num := 1
	changeNum(&num) // passing by reference, passing the address of num
	println("After changeNum:", num) // prints 5, original num is changed

}
