package  main

import "fmt"

func setVal(xPtr *int)  {
	*xPtr = 10
}

func main()  {
	xPtr := new(int)
	setVal(xPtr)
	fmt.Println(*xPtr)
}