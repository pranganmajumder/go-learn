package main

import (
	"fmt"
)

func main() {
	s := make([]string, 3) // this is empty slice
	fmt.Println("emp: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set : ", s)
	fmt.Println("get: ", s[2])

	fmt.Println("len = ", len(s))

	s = append(s, "d")
	fmt.Println("after append  : ", s)

	s = append(s, "e", "f")
	fmt.Println("after append  ", s)

	// another new slice
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("after copy s into c  := ", c)

	l := s[2:5] // from index 2 to (but excluding) s[5]
	fmt.Println("sl1: ", l)

	l = s[:4] //equivalent to s[0:4]
	fmt.Println("sl2 : ", l)

	l = s[2:] // equivalent to s[2:5]
	fmt.Println("sl3 : ", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl", t)

	twoD := make([][]int, 3)  // initialize with row
	for i:=0; i<3; i++{
		innerLen := i+1
		twoD[i] = make([]int, innerLen) // redeclaration twoD
		for j:=0; j<innerLen ; j++{
			twoD[i][j] = i+j
		}
	}
	printTwoD(twoD)

}

func  printTwoD(sl[][] int )  {
	fmt.Println(sl)
}

