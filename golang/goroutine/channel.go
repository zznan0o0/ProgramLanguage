package main

import (
	"fmt"
	"time"
)

func main() {
	a := make(chan int)
	d := 0
	// var a chan int = {0}
	go func(a chan int) {
		a <- 1
		d = 1
	}(a)
	go func(a chan int) {
		a <- 2
		d = 2
	}(a)

	b, c := <-a, <-a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	time.Sleep(1 * time.Second)
	fmt.Println(d)

}
