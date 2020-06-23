package main

import "fmt"

type person struct {
    name string
}

func changeMe1(p *person) {
    *p = person {
        name : "Anthony Hein",
    }
}

func changeMe2(p *person) {
    (*p).name = "Nicholas Hein"
    // p.name = "Nicholas Hein" ALSO OK
}

func main () {
    p := person {
        name : "Michael Hein",
    }
    changeMe1(&p)
    fmt.Println(p)
    changeMe2(&p)
    fmt.Println(p)
}
