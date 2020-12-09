package main

import "fmt"

type Animal struct {
	Name string
	Origin string
}

type Bird struct {
	Animal        // inhereting the Animal Property
	SpeedKPH float32
	CanFly bool
}

func main()  {
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)

	// same type of declaration as before
	//c:= Bird{
	//	Animal: Animal{
	//		Name:   "Doyel",
	//		Origin: "Bangladesh",
	//	},
	//	SpeedKPH: 20,
	//	CanFly:   true,
	//}
	//fmt.Println(c)




}