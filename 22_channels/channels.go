package main

import (
	"fmt"
	// "time"
	// "time"
)

// multiple goroutines should use waitgroups but if one then channel use only

// "math/rand"
// "time"
// "time"

// communication with goroutines

// func processNum(numChan chan int) {
// 	for num := range numChan {
// 		println("Processing number:", num)
// 		time.Sleep(time.Second) // simulate some processing time
// 	}

// }

// func sum(result chan int,num1 int ,num2 int) {
// 	sum := num1 + num2
// 	result <- sum // send the result to the channel
// }

// func task(done chan bool) {

// 	defer func() { done <- true }() // signal that the task is done
// 	println("Task is running")

// }
// <- receive only
// func emailSender(emailChan <-chan string, done chan<- bool) {
// 	defer func() { done <- true }() // signal that the email sending is done
// 	for email := range emailChan {
// 		println("Sending email to:", email)
// 		time.Sleep(time.Second)
// 	}
// }

func main() {
	// messageChan := make(chan string)

	// messageChan <- "Hello, Goroutine!" //blocking

	//  message := <-messageChan // receive message from channel

	//  println("Received message:", message)

	// numChan := make(chan int)

	// go processNum(numChan)

	// for {

	// 	numChan <- rand.Intn(100) // send a random number to the channel
	// }

	// time.Sleep(time.Second * 2) // wait for goroutine to process the number

	// result := make(chan int)
	// go sum(result, 5, 10) // start a goroutine to calculate the sum

	// res := <-result // receive the result from the channel
	// println("Sum is:", res)

	// done := make(chan bool)

	// go task(done)

	// <-done // wait for the task to complete

	// emailChan := make(chan string, 100)

	// done := make(chan bool)

	// go emailSender(emailChan, done) // start the emailSender goroutine
	// // emailChan <- "1@eee.com:example.com"
	// // emailChan <- "2@eee.com:example.com"

	// // fmt.Println(<-emailChan) // receive the first email from the channel
	// // fmt.Println(<-emailChan) // receive the second email from the channel
	// for i:=0;i< 5;i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com",i)
	// }

	// fmt.Println("done")

	// close(emailChan) // close the channel to signal that no more emails will be sent

	// <- done // wait for the emailSender goroutine to finish

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 42 // send an integer to chan1
	}()
	go func() {
		chan2 <- "ping" // send an string to chan2
	}()

	for i:=0; i<2;i++{
		select {
		case chan1Val :=<- chan1:
			fmt.Println("Received from chan1:", chan1Val)
		case chan2Val :=<- chan2:
			fmt.Println("Received from chan2:", chan2Val)
		}
	}
}
