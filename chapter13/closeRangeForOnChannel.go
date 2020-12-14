package main

import "fmt"

func producer(ch chan int)  {
	for i:=0 ; i<10 ; i++{
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan  int)
	go producer(ch)

	for{  // infinite loop
		val, ok := <-ch
		if ok == false{
			break
		}
		fmt.Println("Received val    ", val , ok)
	}

	// equivalent for range

	//for v := range ch{
	//	fmt.Println("received value of channel   ", v )
	//}
}