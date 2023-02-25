package main

import (
	"fmt"
	"log"
)

type student struct {
	name string
}

var data map[int]student

type postgres struct {
	con string
}
type db interface {
	Read(id int) (student, error)
	Delete(id int) error
}

func (p postgres) Read(id int) (student, error) {
	s, ok := data[id]
	if !ok {
		return student{}, fmt.Errorf("record not found in postgres")
	}
	return s, nil
}
func (p postgres) Delete(id int) error {
	delete(data, id)
	return nil
}

type redis struct {
	con string
}

func (r redis) Read(id int) (student, error) {
	s, ok := data[id]
	if !ok {
		return student{}, fmt.Errorf("record not found in redis")
	}
	return s, nil
}

func (r redis) Delete(id int) error {
	delete(data, id)
	return nil
}

func ReadFromDb(d db, id int) {
	s, err := d.Read(id)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%T\n", d)
	fmt.Println(s)
}

func main() {
	data = make(map[int]student)
	data[100] = student{name: "John"}
	r := redis{con: "localhost:redis"}

	p := postgres{con: "root:localhost:postgres"}
	ReadFromDb(p, 100)
	ReadFromDb(r, 100)
}
