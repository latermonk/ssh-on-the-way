package main

import "fmt"

type myint int

//乘2
func (p *myint) mydouble() int {
	*p = *p * 2
	return 0
}

//平方
func (p myint) mysquare() int {
	p = p * p
	fmt.Println("mysquare p = ", p)
	return 0
}

func main() {

	var i myint = 2

	i.mydouble()
	fmt.Println("i = ", i)


	i.mysquare()
	fmt.Println("i = ", i)

}
