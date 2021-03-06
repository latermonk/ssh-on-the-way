// Go语言函数声明语法：函数名之前括号中的内容
// https://blog.csdn.net/xuq09/article/details/88627778


package main

import "fmt"

type Mutatable struct {
	a int
	b int
}

func (m Mutatable) StayTheSame() {
	m.a = 5
	m.b = 7
}

func (m *Mutatable) Mutate() {
	m.a = 5
	m.b = 7
}

func main() {

	m := &Mutatable{0, 0}
	fmt.Println(m)

	m.StayTheSame()
	fmt.Println(m)

	m.Mutate()
	fmt.Println(m)
}
