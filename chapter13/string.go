package  main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(
			// true
			strings.Contains("test", "es"),

			// 2
			strings.Count("test", "t") ,

			// true
			strings.HasPrefix("test", "te"),

			// true
			strings.HasSuffix("test", "st"),

			// 1
			strings.Index("test", "e"),

			// a-b
			strings.Join([]string{"Prangan", "Majumder"} , "-"),

			// == "aaaaa"
			strings.Repeat("a", 5) ,


			// "prbngbn"
			strings.Replace("prangan", "a" , "b" , 2) , // can use 3 or more instead of 2

			// []string{"a", "b", "c", "d", "e"}
			strings.Split("a-b-c-d-e", "-"),

			// "test"
			strings.ToLower("Test"),

			// "TEST"
			strings.ToUpper("Test"),

		)

		// "97 98 100"
		arr := []byte("abd")
		fmt.Println(arr)

		str1 := string([]byte{'t' , 'e' , 's' , 't'})
		str2 := string([]byte{97, 98 , 99})
		fmt.Println(str1 , str2)

		fmt.Printf("Type of str1 & str2  :  %T  %T\n" , str1 , str2)
}