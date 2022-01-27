package main

import "fmt"

type I interface {
	f(num int) string
}

type II struct {
	n1 int
	s1 string
}

func (_this II) f(num int) string {
	return fmt.Sprintf("%s:%d", _this.s1, num+_this.n1)
}

func execI(i I) {
	fmt.Println(i.f(10))
}

func main() {
	ii := II{9, "f_result"}
	execI(ii)
}
