package mypackage

import (
	"fmt"
)

var I int

func Init() {
	I = 0
	fmt.Println("Call mypackage init1")
}

func Init() {
	I = 1
	fmt.Println("Call mypackage init2")
}