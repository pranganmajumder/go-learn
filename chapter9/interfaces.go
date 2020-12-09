package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float32
}

type rect struct {
	height float32
	weidth float32
}

type shape interface {
	area() float32
	//pow() float32
}

func (c circle) area() float32 { // dont use pointer c * circle
	return math.Pi * c.radius * c.radius
}

func (r rect) area() float32  { // dont use pointer
	return r.height * r.weidth
}

func main()  {
	c1 := circle{4.5}
	r1 := rect{5,7}

	//fmt.Println("circle area  " , c1.area())
	//fmt.Println("rectangle area " , r1.area())

	shapes := []shape {c1 , r1}

	for _, shape := range shapes{
		fmt.Println(shape.area())
	}




}