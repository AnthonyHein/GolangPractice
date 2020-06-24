package main

import (
    "fmt"
)

type customErr struct {
    text string
}

func (c customErr) Error() string {
    return c.text
}

func foo(e error) {
    fmt.Println(e.Error())
}

func bar(e error) {
    fmt.Println(e.(customErr).text) // assertion
}

func main () {
    foo(customErr {
        text : "3RR0R",
    })

    bar(customErr {
        text : "3RR0R",
    })
}
