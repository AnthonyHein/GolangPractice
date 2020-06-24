package main

import (
    "fmt"
    "sync"
)

func main () {
    c := make(chan int)

    go func () {
        var wg sync.WaitGroup
        wg.Add(10)
        for i := 0; i < 10; i++ {
            go func () {
                for j := 0; j < 10; j++ {
                    c <- j
                }
                wg.Done()
            }()
        }
        wg.Wait()
        close(c)
    }()

    for {
        v, ok := <-c
        if !ok {
            break
        }
        fmt.Print(v, " ")
    }

    fmt.Println()

    c2 := make(chan int)

    for i := 0; i < 10; i++ {
        go func () {
            for j := 0; j < 10; j++ {
                c2 <- j
            }
        }()
    }

    for k := 0; k < 100; k++ {
        fmt.Println(<-c2)
    }

    close(c2)

    fmt.Println()
}
