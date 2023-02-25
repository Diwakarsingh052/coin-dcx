package main

import "fmt"

type money int // underlying type is int

type student struct { // student is a new data type
	name  string // fields
	age   int    //struct fields would be set to there default if value is not provided
	marks []float64
}

func main() {
	//var dollar money = 100
	//var i int = int(dollar)
	//fmt.Println(dollar)
	var s student // initialize all the fields with its zero value
	s.name = "John"
	s.marks = []float64{10, 20, 30, 50}

	fmt.Println(s)
	fmt.Printf("%+v\n", s) // %+v to print fields and data
	fmt.Printf("%#v\n", s)

	for _, v := range s.marks {
		fmt.Println(v)
	}

	n := s.name
	//string, string, main.student
	fmt.Printf("%T\n", n)
	fmt.Printf("%T\n", s.name)
	fmt.Printf("%T\n", s)

}
