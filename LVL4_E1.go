package main

import "fmt"

func main() {
    a := [5]int{10,20,30,40,50}
    for _, v := range a {
        fmt.Println(v)
    }
    fmt.Printf("%T\n", a)
}
