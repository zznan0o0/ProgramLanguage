package main

import "fmt"

func temp[N int | float64](n1 N, n2 N) N {
	return n1 + n2
}

func main() {
	fmt.Println(temp(1, 2))

	fmt.Println(temp(1.1, 1.2))
}
