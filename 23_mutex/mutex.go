package main

import "sync"

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()  
	p.mu.Lock()     // lock the mutex to ensure safe access to views
	p.views++
	 // unlock the mutex after updating views
}

func main() {

	var wg sync.WaitGroup
	myPost := post{views: 0}
	// myPost.inc()
	// myPost.inc()
	for i := 0; i < 100; i++ {
		wg.Add(1)          // increment the WaitGroup counter
		go myPost.inc(&wg) // increment the views
	}

	wg.Wait()

	println("Post views:", myPost.views)
}
