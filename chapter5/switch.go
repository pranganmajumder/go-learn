package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Scanf("%d", &i)
	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	case 6:
		fmt.Println("Six")
	case 7:
		fmt.Println("Seven")
	case 8:
		fmt.Println("Eight")
	case 9:
		fmt.Println("Nine")
	case 10:
		fmt.Println("Ten")
	default:
		fmt.Println("Unknown number is detected")
	}
}
