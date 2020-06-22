package main

import "fmt"

type person struct {
    fname string
    lname string
    flavors []string
}

func main() {
    p1 := person{
        fname : "Anthony",
        lname : "Hein",
        flavors : []string{"Chocolate", "Mint"},
    }

    p2 := person{
        fname : "Briana",
        lname : "Macedo",
        flavors : []string{"Vanilla", "Cookie Dough"},
    }

    fmt.Println(p1)
    fmt.Println(p2)
}
