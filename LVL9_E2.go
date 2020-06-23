package main

import (
    "fmt"
)

type person struct {
    first string
    last string
    age int
}

func (p *person) speak() {
    fmt.Println(p.first, p.last, p.age)
}

type human interface {
    speak()
}

func saySomething(h human) {
    h.speak()
}

func main () {
    p := person {
        first : "Ant",
        last : "Hein",
        age : 20,
    }

    saySomething(&p)    // Part of interface.
    p.speak()           // Works.
    // saySomething(p)  // Not part of interface.

}
