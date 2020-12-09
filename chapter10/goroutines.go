package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, " : ", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	// go f(0)

	for i := 0; i < 5; i++ {
		go f(i) // it will call randomly every f(i) but will call every function once
	}
	var input string
	fmt.Scanln(&input)
}
