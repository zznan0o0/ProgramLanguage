package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.Atoi(""))

	var e interface{}
	e = 10

	var str string = strconv.Itoa(e.(int))

	fmt.Println(str + "1")
	//e.(type) 不能在switch 以外的地方用
	switch v := e.(type) {
	case int:
		fmt.Println("整型", v)
		var s int
		s = v
		fmt.Println(s)
	case string:
		fmt.Println("字符串", v)
	}

}
