package main

import "fmt"

func main() {
    m := map[string][]string{
        "Hein" : {"Programming", "Frisbee"},
        "Macedo" : {"Doctor", "Reading"},
    }
    for k, v := range m {
        fmt.Println(k)
        for i, val := range v {
            fmt.Printf("Index %d: %s\n", i, val)
        }
    }
}
