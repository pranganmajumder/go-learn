package main

import (
	"fmt"
	"os"
)

func main()  {
	file, err := os.Open("test.txt")
	if err != nil{
		// handle the error here
		fmt.Println("errorrrr 1 ")
		return
	}
	fmt.Println("No error occurred for opening file")
	defer file.Close()



	// get the file size
	stat, err := file.Stat()
	if err != nil{
		// handle the error here
		fmt.Println("errorrrrrr 2")
		return
	}
	fmt.Println("file found and size of chunk : ", stat.Size())


	// read the file
	bs := make([]byte , stat.Size())
	_, err = file.Read(bs)
	if err != nil{
		// handles the error here
		fmt.Println("errorrrrrrrr 3")
		return
	}

	str := string( bs)
	fmt.Println("the data of the file : ", str)
}