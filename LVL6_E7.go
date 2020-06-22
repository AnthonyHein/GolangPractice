package main

import "fmt"

func main () {
    x := func (x int) bool {
        return x % 2 != 0
    }
    fmt.Println(x(7))
}
