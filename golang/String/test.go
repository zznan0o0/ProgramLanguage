package main

import (
	"fmt"
)

func main() {
	a := "a"
	b := "a"

	fmt.Println(&a)
	fmt.Println(&b)

}

func TestLen() {
	str1 := "hello，大哥"
	str2 := "hello, dg"
	fmt.Println(len(str1))
	fmt.Println(len(str2))
}
