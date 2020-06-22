package main

import "fmt"

var x int

func main() {
    defer foo()
    bar()
}

func foo() {
    x *= 2
    fmt.Println("foo", x)
}

func bar() {
    x += 1
    fmt.Println("bar", x)
}
