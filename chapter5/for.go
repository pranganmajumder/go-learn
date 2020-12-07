package main

import (
	"fmt"
)

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//second way
	oddEven()
	divisibilityCheck()
}

// to see another for please go to the file array.go

func oddEven() {
	fmt.Println("second way of printing")
	for i := 1; i <= 10; i++ {
		if i&1 == 1 {
			fmt.Println(i, "is odd")
		} else {
			fmt.Println(i, "is even ")
		}
	}
}

func divisibilityCheck() {
	fmt.Println("divisibility Checking  ")
	for i := 1; i <= 20; i++ {
		if i%4 == 0 {
			fmt.Println(i, " is Divisible by 4")
		} else if i%3 == 0 {
			fmt.Println(i, " is divisible by 3")
		} else if i%2 == 0 {
			fmt.Println(i, "is divisible by 2 ")
		}
	}
}
