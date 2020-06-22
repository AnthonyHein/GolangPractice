package main

import "fmt"

func main() {
    x := make([]string, 3, 8)
    x = []string{"Harvard", "Princeton", "Yale"}
    fmt.Println(x)
    fmt.Println(len(x))
    fmt.Println(cap(x))

    for i := 0; i < len(x); i++ {
        fmt.Println(x[i])
    }
}
