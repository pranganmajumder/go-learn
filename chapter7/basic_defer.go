
package main

import (
	"fmt"
)

func printA(a int)  {
	fmt.Println("Value of a in deferred function  is  ", a)
}

func main()  {

	// arguments

	a := 5
	defer printA(a)

	a = 10
	fmt.Println("value of a before deferred function call is  =  ", a)

}