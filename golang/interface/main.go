package main

import (
    "fmt"
    "errors"
)

type Slice []interface{}

func NewSlice() Slice {
    return make(Slice, 0)
}
 
func (this* Slice) Add(elem interface{}) error {
    for _, v := range *this {
        if v == elem {
            fmt.Printf("Slice:Add elem: %v already exist\n", elem)
            return errors.New("Slice:Add elem: %v already exist\n")
        }
    }
    *this = append(*this, elem)
    fmt.Printf("Slice:Add elem: %v succ\n", elem)
    return nil
}
 
func (this* Slice) Remove(elem interface{}) error {
    found := false
    for i, v := range *this {
        if v == elem {
            if i == len(*this) - 1 {
                *this = (*this)[:i]
 
            } else {
                *this = append((*this)[:i], (*this)[i+1:]...)
            }
            found = true
            break
        }
    }
    if !found {
        fmt.Printf("Slice:Remove elem: %v not exist\n", elem)
        return errors.New("Slice:Remove elem: %v not exist\n")
    }
    fmt.Printf("Slice:Remove elem: %v succ\n", elem)
    return nil
}

func main(){
    arr := NewSlice()
    arr.Add("1")
    arr.Add(1)
    arr.Remove(1)
    fmt.Println(arr)

    var a animal
    var c cat
    a=c
    // var c Cat
    // var d Dog
    a.printInfo()
    // fmt.Println(c.ReturnArg(123))

}

type Animal interface{
    ReturnArg(arg string) string 
}

type Cat int
type Dog []int

// func ReturnArg (arg string) string{
//     return arg
// }

func (c Cat) ReturnArg (arg int) int{
    return arg
}

func (d Dog) ReturnArg (arg []int) []int{
    return arg
}


type animal interface {
    printInfo()
}

type cat int
type dog int

func (c cat) printInfo(){
    fmt.Println(c)
}

func (d dog) printInfo(){
    fmt.Println("a dog")
}


