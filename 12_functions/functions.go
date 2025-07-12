package main

// func add(a int, b int) int {
// 	return a + b
// }
// for same type input parameters we can use below syntax
// func add(a , b int) int {
// 	return a + b
// }'

// used underscore for placeholder if not used to avoid compiler error
// func getLanguage() (string, string, bool) {
// 	return "Go", "Golang", true
// }

// func processIt( fn func(a int)int){
// 	fn(5)
// }

func processIt() func (a int) int {
	return func(a int) int {
		return a * 2
	}
}
func main() {
	// result := add(5, 10)
	// println("The sum is:", result) // Output: The sum is: 15
	// lang1, lang2, test := getLanguage()
	// println("Language:", lang1, "Alias:", lang2, "bool:", test) // Output: Language: Go Alias
	// lang1, lang2, _ := getLanguage()
	// println("Language:", lang1, "Alias:", lang2) // Output: Language: Go Alias

	// fn := func(a int) int {
	// 	return 2
	// }
	// processIt(fn)

	fn := processIt()
	fn(5)
}
