package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	//panic("i need to panic")
	//defer recoverFunc() // func where we recover from panic can't continue to exec and it has to stop
	abc()
	xyz()
}

func abc() {
	defer recoverFunc()
	// recover func recovers from the panic and stops its further propagation
	var i []int
	i[10] = 1000 // panic situation
	fmt.Println(i)
}

func xyz() {
	fmt.Println("I am a normal func ")
}

func recoverFunc() {
	r := recover() // nil means no panic // otherwise r would be the msg of the panic
	if r != nil {
		fmt.Println("recovered from the panic", r)
		fmt.Println(string(debug.Stack()))
	}
}

func search(name string) {
	if name == "" {
		panic("name is not found") // not a good idea , panic should be used when something critical is not working
	}
}
