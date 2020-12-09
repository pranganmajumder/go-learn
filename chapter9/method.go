package main

import (
	"fmt"
)

type Student struct {
	name string
	grade []int
	age int
}

func (s Student) getName() string {
	return s.name
}

func (s *Student) setAge(age int)  {   // using pointer will never create a copy & value can be change using pointer
	s.age = age
}

func (s *Student) getAverage() float32 {
	total := 0
	for _, val := range s.grade{
		total+=val
	}
	return float32(total) / float32(len(s.grade))
}

func main()  {
	s1 := Student{"prangan" , []int{40, 72, 55, 67, 23}, 24}
	fmt.Println(s1.name)


	fmt.Println("before reseting age s1 =  " , s1)
	s1.setAge(50)
	fmt.Println("after reseting age is s1  = " , s1)



	// get average
	avarage1 := s1.getAverage()
	fmt.Println(avarage1)



	s2 := Student{"Tusher Imran" , []int{2, 44, 23} , 20}
	average2 := s2.getAverage()
	fmt.Println(average2)

}