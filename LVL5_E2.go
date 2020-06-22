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

    m := map[string]person{
        "Hein" : p1,
        "Macedo" : p2,
    }

    for _, v := range m {
        fmt.Println(v.fname, v.lname)
        for i, val := range v.flavors {
            fmt.Println(i, val)
        }
    }
}
