package main

import (
	"fmt"
)

type ArrayMapClass interface{
	Print()
}


func Printt(arr1 ArrayMapClass, arr2 ArrayMapClass){
	arr1.Print()
	arr2.Print()
}
type Arr1 []string

func (_this Arr1) Print(){
	fmt.Println(_this[0])
}

type Arr2 []int

func (_this Arr2) Print(){
	fmt.Println(_this[0])
}

func main(){
	arr1 := Arr2{1}
	arr2 := Arr2{2}

	Printt(arr1, arr2)

	arr3 := Arr1{"3"}
	arr4 := Arr1{"4"}
	Printt(arr3, arr4)

	//断言
	var a interface{}
	a = "123"
	value, ok := a.(string)
	if !ok {
		fmt.Println("It's not ok for type string")
		return
	}
	fmt.Println("The value is ", value)

	var t interface{}
	t = 1
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T", t)       // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t)             // t has type bool
	case int:
		fmt.Printf("integer %d\n", t)             // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	}

}



type Sortable interface {
	Len() int
	Less(int, int) bool
	Swap(int, int)
}
 
func bubbleSort(array Sortable) {
	for i := 0; i < array.Len(); i++ {
		for j := 0; j < array.Len()-i-1; j++ {
			if array.Less(j+1, j) {
				array.Swap(j, j+1)
			}
		}
	}
}
 

//实现接口的整型切片
type IntArr []int
 
func (array IntArr) Len() int {
	return len(array)
}
 
func (array IntArr) Less(i int, j int) bool {
	return array[i] < array[j]
}
 
func (array IntArr) Swap(i int, j int) {
	array[i], array[j] = array[j], array[i]
}
 
//实现接口的字符串，按照长度排序
type StringArr []string
 
func (array StringArr) Len() int {
	return len(array)
}
 
func (array StringArr) Less(i int, j int) bool {
	return len(array[i]) < len(array[j])
}
 
func (array StringArr) Swap(i int, j int) {
	array[i], array[j] = array[j], array[i]
}



