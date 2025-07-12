package main

import (
	"fmt"
	"time"
)

// func task(id int) {
// 	println("Task", id, "is running")
// }

func main() {

	for i := 0; i < 10; i++ {
		// go task(i)

		go func (i int) {
			fmt.Println("Task", i, "is running")
		}(i)
	}

	time.Sleep(time.Second * 2) // wait for goroutines to finish
	println("All tasks completed")
}