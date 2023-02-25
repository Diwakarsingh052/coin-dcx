package main

import "fmt"

//type Writer interface {
//	Write(p []byte) (n int, err error)
//}

type expense interface {
	calc() int
}

type perms struct {
	basicPay int
	bonus    int
}

func (p perms) calc() int {
	return p.basicPay + p.bonus
}

type contract struct {
	pay int
}

func (c contract) calc() int {
	return c.pay
}

func TotalExpense(e ...expense) {
	sum := 0
	for _, v := range e {
		sum = sum + v.calc()
	}
	fmt.Println(sum)
}

func main() {

	p := perms{
		basicPay: 100,
		bonus:    100,
	}

	p1 := perms{
		basicPay: 1000,
		bonus:    10,
	}

	c := contract{
		pay: 1000,
	}
	c1 := contract{
		pay: 900,
	}

	TotalExpense(p, p1, c, c1)

}
