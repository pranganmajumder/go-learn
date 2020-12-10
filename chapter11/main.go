package main

import (
	"fmt"
	"github.com/pranganmajumder/learn-go/chapter11/math"
)

func main()  {
	xs := []float32{1,4,5,2, 1}
	avg := math.Average(xs...)
	fmt.Println(avg)
}