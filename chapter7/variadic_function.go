package main

import (
	"fmt"
)

func variadic1(sl1 ... int) int {
	total := 0
	for _, val := range sl1{
		total+=val;
	}
	return  total
}

func multiVariadic2(slice ... [][]int) int { 	// pass twoDimensional slice into a variadic function
	var total int = 0
	for _, row:= range slice{
		for i:= 0 ; i<len(row) ; i++{
			for j:= 0 ; j<len(row[i]) ; j++{
				total += row[i][j]
			}
		}
	}
	return total
}

func main()  {

	fmt.Println("Enter your 3 number")

	sl1 := make([]int , 3)
	for i:= 0 ; i<3 ; i++{
		fmt.Scanf("%d", &sl1[i])
	}
	res := variadic1(sl1...)             // send slice
	fmt.Println(res)

	fmt.Println("multidimentional slice ")
	sl2 := make([][]int, 3)
	for i:= 0 ; i<3 ; i++{
		sl2[i] = make([]int , 2)
		for j:= 0 ; j<2 ; j++{
			fmt.Scanf("%d", &sl2[i][j])
		}
	}

	// pass twoDimensional slice into a variadic function
	var res2 int = multiVariadic2(sl2)
	fmt.Println("sum is  " , res2)
}