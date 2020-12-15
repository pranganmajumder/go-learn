

// practice from https://github.com/spf13/pflag

package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
)

func main() {
	//declares an integer flag, -flagname, stored in the pointer ip , with type *int
	var ip *int = flag.Int("flagname1", 1 , "help message for flagname")



	// bind the flag to a variable using the var() function
	var flagvar int
	func (){
		flag.IntVar(&flagvar, "flagname2" , 2 , "message for flagvar")
	}()

	flag.Parse()



	fmt.Println("ip : ",*ip)
	fmt.Println("flagvar : " , flagvar)
}

