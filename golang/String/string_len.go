package main

import "fmt"

func main() {
	str1 := "hello，大哥"
	str2 := "hello, dg"
	fmt.Println(len(str1))
	fmt.Println(len(str2))
	fmt.Println(fmt.Printf("%014s", str1))
	fmt.Println(fmt.Printf("%014s", str2))
}
