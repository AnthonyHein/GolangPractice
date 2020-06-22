package main

import "fmt"

func main() {
    a := 2000
    for {
        if a > 2020 {
            break
        }
        fmt.Println(a)
        a++
    }
}
