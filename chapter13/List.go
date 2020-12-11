package main

import (
	"container/list"
	"fmt"
)

func main() {
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	for v:=x.Front() ; v!=nil ; v=v.Next(){
		fmt.Println(v.Value)
	}
}