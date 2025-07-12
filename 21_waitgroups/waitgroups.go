package main

import (
	// "fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done() // decrement the WaitGroup counter when the goroutine completes  ,defer runs after its complete
	println("Task", id, "is running")
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		// go task(i)
		wg.Add(1) // increment the WaitGroup counter
		go task(i, &wg)
	}

	wg.Wait() // wait for all goroutines to finish

}
