package  main

import (
	"fmt"
)

type Doctor struct {
	Number int
	ActorName string
	Episodes []string
	Companions []string
}

func  main()  {
	// syntex
	aDoctor := Doctor{
		Number: 3,
		ActorName: "Heisenberg",
		Companions: []string{
			"jessy Pinkman",
			"Skyler",
			"P Scofield",
		},
		Episodes: []string{   // order doesn't matter here
			"episodes 1",
			"episodes 2",
			"episodes 3",
		},
	}
	fmt.Println(aDoctor)



	//// another syntex   if any changes occur in struct blueprint , here need to change
	//// that's why this is discouraged by programmers

	//bDoctor := Doctor{
	//	3,
	//	"Prangan",
	//	[]string{
	//		"jessy Pinkman",
	//		"Skyler",
	//		"P Scofield",
	//	},
	//}
	//fmt.Println(bDoctor)


	// another way of declaring
	newDoctor := struct {
		name string
		degree string
	}{
		name:   "Prangan",
		degree: "MBBS",
	}
	fmt.Println(newDoctor)
	anotherDoctor := newDoctor
	anotherDoctor.name = "Al-Amin Sakib"
	fmt.Println(anotherDoctor.name)


}