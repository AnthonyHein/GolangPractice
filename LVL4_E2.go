package main

import "fmt"

func main() {
    x := []int{10,20,30,40,50, 60, 70, 80, 90, 100}
    for _, v := range x {
        fmt.Println(v)
    }
    fmt.Printf("%T\n", x)
}
