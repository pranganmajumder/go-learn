package main

import (
	"fmt"
	"log"
)

func main() {

	//panic("a problem ")
	//_, err := os.Create("/tmp/file")
	//if err != nil {
	//	panic(err)
	//}

	// according to freeCode camp tutorial
	fmt.Println("start")
	panicker()
	fmt.Println("end")

}

func panicker()  {
	fmt.Println("about to start panicker")

	defer func() {
		if err := recover(); err != nil{
			log.Println("Errorrrr : ", err)
			panic(err)
		}
	}()

	panic("something bad happened")
	fmt.Println("finish panicker function")
}
