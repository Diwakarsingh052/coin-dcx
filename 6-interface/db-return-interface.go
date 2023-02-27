package main

import (
	"fmt"
)

type emp struct {
	name string
}

var details map[int]student

type mysql struct {
	con string
}
type EmpService interface {
	Read(id int) (student, error)
	Delete(id int) error
}

func NewEmpService(e EmpService) EmpService {
	return &mysql{con: "something"}
}

func (p mysql) Read(id int) (student, error) {
	s, ok := data[id]
	if !ok {
		return student{}, fmt.Errorf("record not found in postgres")
	}
	return s, nil
}
func (p mysql) Delete(id int) error {
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

func main() {
	data = make(map[int]student)
	data[100] = student{name: "John"}
	r := redis{con: "localhost:redis"}

	m := mysql{con: "root:localhost:postgres"}
	es := NewEmpService(m)
	ReadFromDb(p, 100)
	ReadFromDb(r, 100)
}
