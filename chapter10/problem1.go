//Lets write one more program to understand channels better. This program will print the sum of the squares and cubes of the individual digits of a number.
//
//For example if 123 is the input, then this program will calculate the output as
//
//squares = (1 * 1) + (2 * 2) + (3 * 3)
//cubes = (1 * 1 * 1) + (2 * 2 * 2) + (3 * 3 * 3)
//output = squares + cubes = 50
//
//We will structure the program such that the squares are calculated in a separate Goroutine, cubes in another Goroutine and the final summation happens in the main Goroutine.


package main

import "fmt"

func calsSquares(number int, sq chan int)  {
	sum := 0
	for number != 0 {
		rem := number % 10
		number /= 10
		sum += rem * rem
	}
	sq <- sum     // write data to channel sq
}

func calcCubes(number int , cu chan  int) {
	sum := 0
	for number != 0 {
		rem := number % 10
		number /= 10
		sum += rem * rem * rem
	}
	cu <- sum   // write data to channel cu
}

func main() {
	number := 589
	sq := make(chan  int)   // create a channel of type int
	cu := make(chan int)

	go calsSquares(number , sq)
	go calcCubes(number , cu)

	sqares , cube := <- sq , <- cu     // read data from channel sq , cu and store into sqares , cube
	fmt.Println("Final Output   " , sqares + cube)
}
// problem link
// https://golangbot.com/channels/