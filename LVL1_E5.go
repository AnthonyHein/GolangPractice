package main

import "fmt"


type ant int
var x ant
var y int

func main() {
    fmt.Println(x)
    fmt.Printf("%T\n", x)
    x = 42
    fmt.Println(x)
    y = int(x)
    fmt.Println(y)
    fmt.Printf("%T\n", y)

}
