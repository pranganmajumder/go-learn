

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

	var ip2 = flag.IntP("ip2Flagname", "f" , 1234 , "helper messege for ip2")
	var flagvar2 bool
	func (){
		flag.BoolVarP(&flagvar2 , "boolname", "b", true , "help message")
	}()
	var flagVal string
	flag.VarP(&flagVal, "varname" , "v" , "help message")
}

