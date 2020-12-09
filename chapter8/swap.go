package  main

import "fmt"

func SWAP(x *int , y *int)  {
	var temp int = *x
	*x = *y
	*y = temp

}

func main()  {
	var x int = 1
	var y int = 2
	SWAP(&x, &y)
	
	fmt.Println("after swapping  x : " , x , "  y : " , y)
}