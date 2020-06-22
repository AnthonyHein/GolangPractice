package main

import "fmt"

type person struct {
    first string
    last string
    age int
}

func (p person) speak() {
    fmt.Println(p.first, p.last, p.age)
}

func main() {
    p := person {
        first : "Anthony",
        last : "Hein",
        age : 20,
    }
    p.speak()
}
