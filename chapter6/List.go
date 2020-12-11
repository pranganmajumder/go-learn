package main

import (
	"container/list"
	"fmt"
)

func main() {
	var x list.List
	x.PushBack(10)
	x.PushBack(20)
	x.PushBack(30)

	for v := x.Front() ; v!=nil ; v = v.Next(){
		fmt.Println(v.Value)
	}
}