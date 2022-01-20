package main

import "fmt"

// 定义结构
type  abc struct {

}


func (i *abc) test() int{
	fmt.Println("***********************")
	return 100
}

func main(){
	// 初始化结构
	aaa := &abc{

	}

	//调用结构的方法
	aaa.test()

}
