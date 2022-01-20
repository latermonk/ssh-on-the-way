package main

import (
	"log"

	"time"
)

func main() {

	ch := make(chan int)

	go task1(ch)
	go task2(ch)

	for i := 0; i < 2; i++ {

		v := <-ch
		log.Println("one task done:", v)
	}

	log.Println("All task done")

}

func task1(c chan int) {

	time.Sleep(1 * time.Second)

	c <- 1

}

func task2(c chan int) {

	time.Sleep(10 * time.Second)

	c <- 2

}
