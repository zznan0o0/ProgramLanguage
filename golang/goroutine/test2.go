package main

import (
	"fmt"
	"time"
)

func main() {
	a := 0

	go func() {
		a = 1
	}()

	fmt.Println(a)
	time.Sleep(1 * time.Second)
	fmt.Println(a)
}
