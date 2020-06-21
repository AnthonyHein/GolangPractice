package main

import "fmt"


type ant int
var x ant

func main() {
    fmt.Println(x)
    fmt.Printf("%T\n", x)
    x = 42
    fmt.Println(x)
}
