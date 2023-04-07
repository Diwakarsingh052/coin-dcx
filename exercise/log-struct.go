package main

import (
	"log"
	"os"
)

func main() {
	//create a struct that store connection to log
	//create var of struct using new func
	//write a method that log on terminal
	l := NewLogging()
	l.print("test data") // date + time + line number
	//fmt.Println(string(debug.Stack()))

}

type Logging struct {
	// connection to log.Logger

}

func NewLogging() *Logging {
	l := log.New(os.Stdout, "sales: ", log.LstdFlags|log.Lshortfile)
	return &Logging{}
}

func (l *Logging) print(data string) {
	//print data
}
