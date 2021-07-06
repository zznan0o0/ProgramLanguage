package main

import "fmt"

func sum(s []int, c chan int) {
		fmt.Println(s)
        sum := 0
        for _, v := range s {
                sum += v
        }
        c <- sum // 把 sum 发送到通道 c
}

func main() {
        s := []int{7, 2, 8, -9, 4, 0}

        c := make(chan int)

		fmt.Println(123456)
        go sum(s[:len(s)/2], c)
        go sum(s[len(s)/2:], c)
        go sum(s[len(s)/3:], c)
        go sum([]int{3}, c)
		fmt.Println(1234567)

        // x, y := <-c, <-c // 从通道 c 中接收
        // 执行顺序不一定的
        x, y, z, a := <-c, <-c, <-c, <-c // 从通道 c 中接收

        // fmt.Println(x, y)
        fmt.Println(x, y, z, a)
}