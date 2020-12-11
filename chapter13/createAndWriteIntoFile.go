package main

import (
	"os"
)

func main() {
	file, err := os.Create("test2.txt")
	if err != nil{
		// handle the error there
		return
	}
	defer file.Close()
	file.WriteString("this is our newly created file test2.txt")


}
