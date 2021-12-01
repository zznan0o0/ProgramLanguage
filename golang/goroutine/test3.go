package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 0; i <= 10; i++ {
		fmt.Println("hello world+" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	go test()
	for i := 0; i <= 10; i++ {
		fmt.Println("hello golannd+" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
