package main

import (
	"fmt"
	"time"
)

func main() {

	var num int

	go func() {
		for {
			fmt.Println(num)
			num += 1
			fmt.Println(num)
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		for {
			fmt.Println(num)
			num += 1
			fmt.Println("end:", num)
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		for {
			fmt.Println(num)
			num += 1
			fmt.Println("end:", num)
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(1 * time.Hour)
}
