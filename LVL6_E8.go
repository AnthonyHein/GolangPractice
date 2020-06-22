package main

import "fmt"

func main () {
    myFunc := inc(3)
    fmt.Println(myFunc())
    fmt.Println(myFunc())
    fmt.Println(myFunc())
}

func inc(rate int) func() int {
    x := 0
    f := func() int {
        x += rate
        return x
    }
    return f
}
