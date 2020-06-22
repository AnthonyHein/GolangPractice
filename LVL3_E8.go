package main

import "fmt"

func main() {
    switch {
        case false:
            fmt.Println("FALSE")
        case 8 < 7:
            fmt.Println("FALSE")
        case true:
            fmt.Println("YES!")
    }
}
