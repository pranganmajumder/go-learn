package main

import (
	"fmt"
)

func makeEvenGenerator() func() int {
	i := 0
	return func() int {
		i += 2
		return i
	}
}

func appendStr() func(string) string {
	h := "Hello"
	g := func(m string) string {
		h = h + " " + m
		return h
	}
	return g
}

func main() {
	nextEven := makeEvenGenerator()
	fmt.Println("initial state  ", nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	nextEven = makeEvenGenerator()
	fmt.Println("back to initial state   ", nextEven(), " \n ") // initial state

	a := appendStr()
	b := appendStr()

	fmt.Println(a("World"))
	fmt.Println(b("Everyone"))

	fmt.Println(a("Gopher"))
	fmt.Println(b("!"))

}
