package main

import "fmt"

type Runner interface {
	Run()
}
type Walker interface {
	Walk()
}

type WalkRunner interface {
	Runner // embedding interface
	Walker
}

type Human struct{ name string }

func (h Human) Walk() {
}
func (h Human) Run() {
}

type Robo struct{ name string }

func (r Robo) Run() {}

func main() {

	h := Human{name: "John"}
	robot := Robo{name: "r1"}

	_, _ = h, robot

	CanWalk(h)
	CanWalk(robot)

}

func CanWalk(r Runner) {
	//type assertion
	i, ok := r.(Human) // check whether a type exists in the interface or not // if it is there than that struct would be returned

	if !ok { // ok would be true if the human struct is present otherwise it would be false
		fmt.Println("you are not human, can't call walk method")
		return
	}
	fmt.Println("calling the walk method")
	i.Walk()
}
