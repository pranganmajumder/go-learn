
// When a function has multiple defer calls,
//they are pushed on to a stack and executed in Last In First Out (LIFO) order.

package main

import (
	"fmt"
)


func main()  {

	// stack of defers
	name := "Prangan"
	fmt.Printf("Original String : %s\n", string(name))

	fmt.Printf("Reversed String is  : ")
	for _, ch := range []rune(name){
		defer fmt.Printf( "%c", ch)
	}
}