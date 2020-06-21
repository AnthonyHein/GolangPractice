package main

import "fmt"

const x = 42
const y uint8 = 255

func main() {
    fmt.Println(x)
    fmt.Println(y)
    fmt.Printf("%T\n", x)
    fmt.Printf("%T\n", y)
}
