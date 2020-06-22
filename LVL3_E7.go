package main

import "fmt"

func main() {
    if x := 23; x == 42 {
        fmt.Println(x)
    } else if x < 42 {
        fmt.Println("Low")
    } else {
        fmt.Println("High")
    }
}
