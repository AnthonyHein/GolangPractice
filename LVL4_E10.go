package main

import "fmt"

func main() {
    m := map[string][]string{
        "Hein" : {"Programming", "Frisbee"},
        "Macedo" : {"Doctor", "Reading"},
    }
    m["Foglia"] = []string{"Research", "Memes"}
    delete(m, "Macedo")
    for k, v := range m {
        fmt.Println(k)
        for i, val := range v {
            fmt.Printf("Index %d: %s\n", i, val)
        }
    }
}
