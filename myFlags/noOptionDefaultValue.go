// resources : https://github.com/spf13/pflag

package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
)



func main() {
	var ip = flag.IntP("flagname", "f" , 1234 , "help message")
	flag.Lookup("flagname").NoOptDefVal = "4321"
	flag.Parse()
	fmt.Println("ip=",*ip)


}