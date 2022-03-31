package main

import "fmt"

func temp[N int | float64](n1 N, n2 N) N {
	return n1 + n2
}

func temp2[T any](n1 T, n2 T, fn func(T, T) T) T {
	return fn(n1, n2)
}

func main() {
	fmt.Println(temp(1, 2))

	fmt.Println(temp(1.1, 1.2))
	fmt.Println(temp2(1, 4, func(t1, t2 int) int { return t1 + t2 }))
}
