package main

import "fmt"

func getSum(pos int, ar [6]int) int {
	if pos == len(ar) {
		return 0
	}
	return ar[pos] + getSum(pos+1, ar)
}

func main() {
	ar := [...]int{2, 4, 6, 7, 3, 6}
	sum := getSum(0, ar)
	fmt.Println("sum is   =  ", sum)
}
