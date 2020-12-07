package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a number :   ")
	var input float32
	fmt.Scan(&input)

	output := input + 7
	fmt.Println(output)

}
