package main

import (
    "fmt"
    "sync"
)

func main () {

    incVal := 0

    gs := 100
    var wg sync.WaitGroup
    wg.Add(gs)
    var mu sync.Mutex

    for i := 0; i < gs; i++ {
        go func () {
            mu.Lock()
            newVal := incVal
            newVal++
            incVal = newVal
            fmt.Println("Int val:", incVal)
            mu.Unlock()
            wg.Done()
        }()
    }
    
    wg.Wait()
    fmt.Println("Val:", incVal)
}
