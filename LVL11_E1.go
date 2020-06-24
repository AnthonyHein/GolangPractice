package main

import (
    "fmt"
    "encoding/json"
    "log"
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

    bs, err := json.Marshal(p1)
    if err != nil {
        log.Println(err)
        log.Println(err.Error())
    }
    fmt.Println(string(bs))
}
