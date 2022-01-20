package main

import "fmt"

// 定义结构
type Abc struct {
}

func (i *Abc) test() int {
	fmt.Println("***********************")
	return 100
}

func main() {
	// 初始化结构
	aaa := &Abc{}

	//调用结构的方法
	aaa.test()

}
