package main

import "fmt"

var map_data map[string]string = map[string]string{}

// func init() {
// 	map_data = map[string]string{}
// }

type aaa struct {
	a string
}

func main() {
	map_data["sadasd"] = "sadasdasd"

	fmt.Println(map_data)

	m := map[string]aaa{}
	m["aaa"] = aaa{"aaaaa"}

}
