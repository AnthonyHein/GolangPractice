package main

import "fmt"

func main() {
    favSport := "Baseball"
    switch favSport {
        case "Football":
                fmt.Println("NO!")
        case "Baseball":
            fmt.Println("NICE!")
    }
}
