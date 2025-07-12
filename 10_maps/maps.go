package main

import "maps"

// import "fmt"

// maps -> hash,object,dict
func main() {
	// 	m := make(map[string]string)

	// 	//setting elements
	// 	m["name"] = "John Doe"
	// m["age"] = "30"
	// 	// getting elements
	// 	println(m["name"],m["ag"]) // prints "John Doe"
	// 	// if key does not exist then it returns zero value
	// 	println(m["address"]) // prints "" (empty string)

	// m :=make(map[string]int)
	// m["a"]=1

	// println(m["age"],len(m)) // prints 0 (zero value for int)

	// delete(m,"a") // delete a key-value pair

	// clear(m) // clear the map

	// m := map[string]int{"a": 1, "b": 2, "c": 3} // map literal
	// println(m["a"]) // prints 1

	// m := map[string]int{"price": 100, "banana": 50, "apple": 80} // map literal

	// _, ok := m["banana"] // get value and ok if key exists
	// if ok {
	// 	println("banana exists with value:", ok) // prints "banana exists with value: 50"
	// } else {
	// 	println("banana does not exist")
	// }
	m := map[string]int{"price": 100, "banana": 50, "apple": 80}

	m2 := map[string]int{"price": 100, "banana": 50, "apple": 80}

	println(maps.Equal(m, m2)) // prints true
}
