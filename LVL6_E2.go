package main

import "fmt"

func main() {
    fmt.Println(foo([]int{1,3,5,7}...))
    fmt.Println(bar([]int{2,4,6,8}))
}

func foo(x ...int) int {
    var sum int
    for _, v := range x {
        sum += v
    }
    return sum
}

func bar(x []int) int {
    var sum int
    for _, v := range x {
        sum += v
    }
    return sum
}
