package main

import (
	"fmt"
	"time"
)

func hello(done chan bool)  {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)

	fmt.Println("hello go routine awake and going to write to done")

	done <- true
}

func main() {

	//// create a channel
	//var a chan int
	//if a == nil{
	//	fmt.Println("channel a is nill , going to define it ")
	//	a = make(chan int)
	//	fmt.Printf("Type of a is %T" , a)
	//
	//	//   data := <- a   // read data from channel a & storing the value to the variable "data"
	//	//   a <- data  	// write data to channel a
	//}



	done := make(chan bool)
	fmt.Println("Main is going to call hello go goroutine")

	go hello(done)
	<- done

	fmt.Println("Main received data")

}