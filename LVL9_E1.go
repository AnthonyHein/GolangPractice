package main

import (
    "fmt"
    "sync"
    "runtime"
)

func main () {
    var wg sync.WaitGroup

    wg.Add(2)

    fmt.Println("Which came first?")
    fmt.Println("routines:", runtime.NumGoroutine())

    go func () {
        fmt.Println("Chicken")
        wg.Done()
    }()

    go func () {
        fmt.Println("Egg")
        wg.Done()
    }()

    fmt.Println("routines:", runtime.NumGoroutine())

    wg.Wait()

    fmt.Println("There is your answer.")


}
