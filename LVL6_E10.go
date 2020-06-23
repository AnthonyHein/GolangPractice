package main

import "fmt"

func main () {
    f := myVar()
    fmt.Println(f())
    fmt.Println(f())
    fmt.Println(f())
}

func myVar() func() int {
    var x int
    return func () int {
        x++
        return x
    }
}
