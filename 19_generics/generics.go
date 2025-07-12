package main

// import "fmt"

// after 1.18 version
// generics - type parameters
// func printSlice(s []int) {
// 	for _, v := range s {
// 		println(v)
// 	}
// }
// func printStringSlice(s []string) {
// 	for _, v := range s {
// 		println(v)
// 	}
// }

// any  = interface{}
// type parameters - generic types
// func printSlice[T any](s []T) {
// 	for _, v := range s {
// 		println(v)
// 	}
// }
// func printSlice[T int | string | bool](s []T) {
// 	for _, v := range s {
// 		println(v)
// 	}
// }

//  LIFO
// type Stack struct {
// 	elements []int
// }
// type Stack[T any] struct {
// 	elements []T
// }

func printSlice[T comparable, V string](s []T , name V) {
	for _, v := range s {
		println(v,name)
	}
}
func main() {

	// nums :=[]int{1, 2, 3, 4, 5}
	// printSlice(nums)

	// nums :=[]string{"one","two"}
	// printStringSlice(nums)

	// nums := []int{1, 2, 3, 4, 5} 
	// vals := []string{"one", "two", "three"}
	// bools := []bool{true, false, true}
	// printSlice(vals)
	// printSlice(bools)
	// printSlice(nums)

	// myStack := Stack{
	// 	elements: []int{1,2,3},
	// }
	// fmt.Println( myStack)
	// myStack := Stack[string]{
	// 	elements: []string{"go","ok"},
	// }
	// fmt.Println( myStack)

	
	// nums := []int{1, 2, 3, 4, 5} 
	vals := []string{"one", "two", "three"}
	// bools := []bool{true, false, true}
	printSlice(vals, "string")
	// printSlice(bools)
	// printSlice(nums)
}