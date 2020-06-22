package main

import "fmt"

type vehicle struct {
    doors int
    color string
}

func main() {
    ss := struct {
        vehicle
        luxury bool
    }{
        vehicle : vehicle{
            doors : 4,
            color : "Beige",
        },
        luxury : true,
    }

    fmt.Println(ss)

}
