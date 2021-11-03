package main

import "fmt"

func main() {
	a := 1
	var b *int = &a
	d := &a
	c := *b
	fmt.Println(d)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(&c)

	a = 2

	fmt.Println(a)
	fmt.Println(*b)
	fmt.Println(c)

}
