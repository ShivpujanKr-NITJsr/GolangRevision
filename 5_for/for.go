package main

// import "fmt"

//for loop only ,no while loop

func main() {
	//while loop
	// i:=1
	// for i <= 3{ // condition
	// 	fmt.Println(i) // body
	// 	i++ // increment
	// }
	// infinite loop
	// for {
	// 	println("Hello World")
	// }

	//classic for loop

	// for i := 1; i <= 5; i++ { // initialization, condition, increment
	// 	if i == 1 {
	// 		continue // skip the rest of the loop body and continue with the next iteration
	// 	}
	// 	println(i) // body
	// 	break      // exit the loop
	// }
	// 1.22 version =range based for loop
	for i := range 3 {
		println(i) // prints 0, 1, 2
	}
}
