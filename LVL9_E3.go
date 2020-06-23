package main

import (
    "fmt"
    "sync"
    "runtime"
)

func main () {

    incVal := 0

    gs := 100
    var wg sync.WaitGroup
    wg.Add(gs)

    for i := 0; i < gs; i++ {
        go func () {
            newVal := incVal
            runtime.Gosched() // yield processor to different function
            newVal++
            incVal = newVal
            fmt.Println("Int val:", incVal)
            wg.Done()
        }()
    }

    fmt.Println("Val:", incVal)
    wg.Wait()
}
