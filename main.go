package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word" , "foo", "a string")

	numbPtr := flag.Int("numb" , 42 , "an integer")
	boolPtr := flag.Bool("fork" , false , "a bool")

	var svarr string
	flag.StringVar(&svarr, "svar" , "bar" , "a string var") // created custom flag

	flag.Parse()

	fmt.Println("word : " , *wordPtr)
	fmt.Println("numb : " , *numbPtr)
	fmt.Println("fork : " , *boolPtr)
	fmt.Println("tail : " , svar)
	fmt.Println("tail: " , flag.Args())

}