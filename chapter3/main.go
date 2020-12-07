package main

import (
	"fmt"
)

func main() {
	fmt.Println("1 + 1 = ", 1.19+1.0)
	fmt.Println(5 * 8)
	fmt.Println("   ", 12%5)

	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")

	// boolean
	fmt.Println("\nBooleans     ======>>>>> ")
	fmt.Println(true && true)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println((true && false) || (false && true) || !(false && false))

}
