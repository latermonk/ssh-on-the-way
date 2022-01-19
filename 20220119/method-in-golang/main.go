// Go基础系列：Go中的方法
//https://www.cnblogs.com/f-ck-need-u/p/9890624.html

// Go中的struct结构类似于面向对象中的类。面向对象中，除了成员变量还有方法。
//Go中也有方法，它是一种特殊的函数，定义于struct之上(与struct关联、绑定)，被称为struct的receiver。


package main

import "fmt"

type changfangxing struct {
	length float64
	width  float64
}

func (c *changfangxing) area() float64 {
	return c.length * c.width
}

func main() {
	c := &changfangxing{
		2.5,
		4.0,
	}
	fmt.Printf("%f\n",c.area() )
}