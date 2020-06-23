package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main () {
    var wg sync.WaitGroup
    var incVal int64

    gs := 100
    wg.Add(gs)

    for i := 0; i < gs; i++ {
        go func () {
            atomic.AddInt64(&incVal, 1)
            fmt.Println("Int val:", atomic.LoadInt64(&incVal))
            wg.Done()
        }()
    }

    wg.Wait()
    fmt.Println("Val:", incVal)
}
