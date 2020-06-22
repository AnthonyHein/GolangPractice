package main

import "fmt"

func main () {
    f := func (x int) int {
        return x * x
    }
    fmt.Println(sum(f))

}

func sum (myFunc func(int) int) int {
    var total int
    for i := 0; i < 5; i++ {
        total += myFunc(i)
    }
    return total
}
