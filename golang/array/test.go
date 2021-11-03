package main

import "fmt"

func main() {
	var a [3]int = [3]int{1, 2, 3}
	// append(a, 4) is not a slice
	var b []int = []int{1, 2, 3, 4}
	b = append(b, 5)
	fmt.Println(a)
	fmt.Println(b)
}
