package main

import "fmt"

func main () {
    x := func (x int) bool {
            return x % 2 == 0
        }(7)
    fmt.Println(x)
}
