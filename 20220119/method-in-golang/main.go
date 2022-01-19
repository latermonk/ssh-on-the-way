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
	fmt.Printf("%f\n",c.area())
}