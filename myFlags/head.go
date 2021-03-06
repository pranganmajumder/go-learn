// resources : https://www.digitalocean.com/community/tutorials/how-to-use-the-flag-package-in-go
// displays the first several lines of a given file:
// run the file : go run head.go -- filename
// ex : go run head.go -- boolean.go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var count int
	flag.IntVar(&count , "n", 5 , "number of lines to read from the file") // value = number of lines want to display
	flag.Parse()

	var in io.Reader
	if filename := flag.Arg(0); filename != ""{
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println("error opening file: ", err)
			os.Exit(1)
		}
		defer f.Close()
		in = f
	}else{
		in = os.Stdin
	}

	buf := bufio.NewScanner(in)

	for i := 0 ; i<count ; i++{
		if !buf.Scan(){
			break
		}
		fmt.Println(buf.Text())
	}
	if err := buf.Err() ; err != nil{
		fmt.Fprintln(os.Stderr , "error reading : err", err)
	}
}