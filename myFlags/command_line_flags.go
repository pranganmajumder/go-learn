// according to https://gobyexample.com/command-line-flags

package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word" , "foo", "a string")

	numbPtr := flag.Int("numb" , 42 , "an integer")
	boolPtr := flag.Bool("fork" , false , "a bool")

	// create custom flag
	var svarr string
	flag.StringVar(&svarr , "svar" , "bar" , "a string var")

	flag.Parse()

	fmt.Println("word : " , *wordPtr)
	fmt.Println("numb : " , *numbPtr)
	fmt.Println("fork : " , *boolPtr)
	fmt.Println("tail : " , svarr)
	fmt.Println("tail: " , flag.Args())

}


//after build the go file
// run the binary file using ./command_line_flags
// to set the value ./command_line_flags -word=Bangladesh -numb=420 -fork=true -svar=hi