package main

import (
	"fmt"
	"log"
	"strconv"
)

func SumString(s, x string) (int, error) { // err must be the last value to be returned
	a, err := strconv.Atoi(s)
	if err != nil {
		//log.Println(err) // avoid the logs // error handling should be done by the caller
		return 0, err // when err happens then set other values to their defaults
	}

	b, err := strconv.Atoi(x)
	if err != nil {
		//log.Println(err) // if you are logging an error then consider the error is handled // don't propagate it further
		return 0, err // when err happens then set other values to their defaults
	}

	sum := a + b
	return sum, nil //success

}

func main() {

	sum, err := SumString("abc", "20")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(sum)

}
