package main

import (
	"fmt"
)

var st string = "prangan"
var i int = 345

// this is another way of declaring global variable
var (
	actorName    string = "Dipika Padukone"
	companion    string = "Ranbir Singh"
	doctorNumber int    = 3
	season       int    = 11
)

// this is the another way of declaring global variable
var (
	g1 = "g1"
	g2 = 2222
	g3 = 34.344
)

func main() {
	// first way
	fmt.Println(" i = ", i)
	var i int
	i = 43
	fmt.Println("i = ", i)

	// second way
	var j float64 = 100
	fmt.Println(j)

	// third way
	k := 199
	fmt.Println("the value of k is  ", k)
	fmt.Printf("value and type of k is %v  %T\n  ", k, k)
	fmt.Printf("%v %T\n", st, st)
	fmt.Println("hello, my name is")

	fmt.Println("Actor name is  ", actorName)

	// Redeclaration of a variable
	var a int = 420
	a = 421
	fmt.Printf("%v %T\n", a, a)

	// typeCasting
	var typeCast float32
	typeCast = float32(a)
	fmt.Printf("%v %T\n", typeCast, typeCast)

	// another way
	var x = "check"
	fmt.Println("print the value x  ", x)
	fmt.Printf("%v  %T\n", x, x)

	// print global variable g1, g2, g3
	g1g2g3()

}

func g1g2g3() {
	fmt.Println("Print the global variable g1, g2 ,g3 ")
	fmt.Println("  g1 = ", g1)
	fmt.Println("  g2  = ", g2)
	fmt.Println("  g3  = ", g3)
}
