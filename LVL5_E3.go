package main

import "fmt"

type vehicle struct {
    doors int
    color string
}

type truck struct {
    vehicle
    fourWheel bool
}

type sedan struct {
    vehicle
    luxury bool
}

func main() {
    st := truck{
        vehicle : vehicle{
            doors : 2,
            color : "Black",
        },
        fourWheel : false,
    }
    ss := sedan{
        vehicle : vehicle{
            doors : 4,
            color : "Beige",
        },
        luxury : true,
    }

    fmt.Println(st, ss)
    fmt.Println(st.doors, ss.luxury)

}
