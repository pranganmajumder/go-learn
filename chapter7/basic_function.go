package main

import "fmt"

func plus(a int , b int )int  {
	return  a+b
}

func plusPlus(a, b, c int) int  {
	return  a+b+c
}


// returning multiple values from a function
func multipleValues() (int,int)  {
	return 5,6
}


func main()  {

	// return single value from function
	res:= plus(1,2)
	fmt.Println(res)
	fmt.Println(plusPlus(1,2,3))


	// return multiple variables from function
	x, err := multipleValues()
	fmt.Println(x, err)
	x,err = multipleValues()
	fmt.Println(err , x)
}


