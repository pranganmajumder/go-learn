package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter five integer number  ")
	var array [5]int
	for i := 0; i < 5; i++ {
		fmt.Scanf("%d", &array[i])
	}
	firstWay(array)
	secondWay(array)

	// shorter syntex for  creating array
	array2 := [5]int{11, 12, 13, 14, 15}
	secondWay(array2)

	// another shorter syntex for creating array
	array3 := [5]int{
		1,
		2,
		3,
		4,
		5,
	}
	secondWay(array3)

	// two dimensional array
	var twoD [3][5]int // row , column
	fmt.Println("Enter your input for two dimensional array 3 row , 5 column")
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			fmt.Scanf("%d", &twoD[i][j])
		}
	}
	printTwoDArray(twoD)

	//another way of creating array
	ar1 := [3]int{12} // declares an array of len 3 output [12, 0, 0]
	ar2 := [...]int{12, 78, 50}
	fmt.Println(ar1)      // 12, 0 , 0
	fmt.Println(len(ar2)) // 3
}

func printTwoDArray(twoD [3][5]int) {
	fmt.Println("Printing Two dimentional array")
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print(twoD[i][j], " ")
		}
		fmt.Println()
	}
}

func firstWay(array [5]int) {

	var cumSum int = 0
	for i := 0; i < len(array); i++ {
		cumSum += array[i]
	}
	var avgSum float32 = float32(cumSum) / 5 // type casting
	fmt.Println("avarage of the array is ", avgSum)
}

func secondWay(array [5]int) {
	var cumSum int = 0
	for _, value := range array {
		cumSum += value
	}
	var avgSum float32 = float32(cumSum) / 5
	fmt.Println("average of the array is ", avgSum)
}
