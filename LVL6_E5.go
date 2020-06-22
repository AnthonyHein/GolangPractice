package main

import "fmt"
import "math"

type square struct {
    edge float64
}

type circle struct {
    radius float64
}

func (s square) area() float64 {
    return s.edge * s.edge
}

func (c circle) area() float64 {
    return c.radius * c.radius * math.Pi
}

type shape interface {
    area() float64
}

func info(s shape) {
    fmt.Println(s.area())
}

func main() {
    si := square {
        edge : 1,
    }
    ci := circle {
        radius : 1,
    }

    info(si)
    info(ci)
}
