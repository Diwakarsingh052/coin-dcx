package main

import "fmt"

type student struct {
	name  string
	email string
}

func (s student) print() {
	fmt.Println(s.name, s.email)
}

func main() {
	var s1 []student // default value is nil
	_ = s1
	s := []student{
		{
			name:  "s1",
			email: "s1@email.com",
		},
		{
			name:  "s2",
			email: "s2@email.com",
		},
	}
	s = append(s, student{
		name:  "s3",
		email: "s3@email.com",
	})

	for _, stu := range s {
		stu.print()
	}
}
