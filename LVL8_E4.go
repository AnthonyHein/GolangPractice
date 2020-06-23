package main

import (
    "fmt"
    "sort"
)

func main () {
    xi := []int{5, 8, 2, 43, 17}
    xs := []string{"random", "rainbow", "delights", "in", "torpedo"}

    fmt.Println(xi)
    fmt.Println(xs)

    sort.Ints(xi)
    sort.Strings(xs)

    fmt.Println(xi)
    fmt.Println(xs)
}
