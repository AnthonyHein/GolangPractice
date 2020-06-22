package main

import "fmt"

func main() {
    xs := [][]string{
        []string{"James", "Bond"},
        []string{"Miss", "Moneypenny"},
    }

    for _, p := range xs {
        for _, v := range p {
            fmt.Print(v, " ")
        }
        fmt.Println()
    }
}
