package main

import "fmt"

func main()  {
	m := make(map[string]int)  // create an empty map
	m["key1"] = 7
	m["key2"] = 8

	fmt.Println("initial map is = " , m)

	v1 := m["key1"]          // assign into others variable

	fmt.Println("v1  = ", v1)
	fmt.Println("len of map = ", len(m))  // check len

	delete(m, "key2")  				// delete a key

	fmt.Println("updated map is   = "  , m)


	val, ok := m["key1"]   // first is value , second is remain or not
	fmt.Println(val, ok)   // output is 7, true

	_ , ok = m["key2"]
	fmt.Println(ok)        // false , because "key2" isn't present here







	// declare and initialize a new map in the same line
	newMap := map[string]int {"foo" :1 , "bar": 2}
	fmt.Println("new map is   =   " , newMap)


	// check if a key remain or not in map
	if val, ok := m["key1"]; ok{ 				// 1st we try to get the value from the map
		fmt.Println("yes  = " , val, ok)    // if successful we run the code inside of the block
	}


	// shorter way of creating map
	shorterMap := map[string]int{
		"key1" : 1,
		"key2" : 2,
		"key3" : 3,
		"key4" : 4,
	}
	fmt.Println("shorterMap =   " , shorterMap)

	multiDimentionalMap()

}

func multiDimentionalMap()  {
	elements := map[string]map[string]string{  // key value
		"H" : map[string]string{
			"name"  : "Hydrogen",
			"state" : "gas" ,
		},
		"He" : map[string]string{
			"name" : "Helium",
			"state" : "gas",
		},
	}

	if val, ok := elements["H"]; ok{
		fmt.Println("multiDimentional  Map   =  " , val["name"] , val["state"])
	}
}