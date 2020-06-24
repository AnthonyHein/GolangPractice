package main

import (
    "fmt"
)

type person struct {
    First string
    Last string
    Sayings []string
}

func main () {
    p1 := person {
        First : "James",
        Last : "Bond",
        Sayings : []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
    }

    bs, err := toJSON(p1)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(string(bs))
}

func toJSON(p person) ([]byte, error) {
    return []byte{}, fmt.Errorf("My first formatted error. %v", 5)
}
