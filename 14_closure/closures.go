package main

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	increment := counter() // create a closure
	println(increment()) // prints 1
	println(increment()) // prints 2
}