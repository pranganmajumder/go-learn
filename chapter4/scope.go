package main

import (
	"fmt"
)

func main() {
	{
		var inner = "inner"
		fmt.Println(inner)
	}
	// fmt.Println(inner) // this line can't be used because variable exists within the nearest curly braces { }
	var inner = "outer scope"
	fmt.Println(inner)
}
