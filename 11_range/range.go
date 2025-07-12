package main

func main() {
	// nums := []int{1, 2, 3, 4, 5}
	// for i, v := range nums {
	// 	println("Index:", i, "Value:", v)
	// }

	// for i :=0 ;i<len(nums); i++ {
	// 	println("Index:", i, "Value:", nums[i])
	// }

	// for i, v := range nums {
	// 	println( i ,"Value:", v)
	// }

	// m :=map[string]int{"a": 1, "b": 2, "c": 3}
	// for k, v := range m {
	// 	println("Key:", k, "Value:", v)
	// }
//byte=255,but unicode can be larger so we use rune,index will be skip one byte in below example
	for i, v := range "Hello" {
		println(v,"Index:", i, "Value:", string(v)) // prints index and character,v=ascii
	}
}