package main

import (
    "fmt"
    "sort"
)

type user struct {
    Last string
    Age int
    Sayings []string
}

type ByAge []user

func (a ByAge) Len() int {
    return len(a)
}

func (a ByAge) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
    return a[i].Age < a[j].Age
}

type ByLast []user

func (a ByLast) Len() int {
    return len(a)
}

func (a ByLast) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}

func (a ByLast) Less(i, j int) bool {
    return a[i].Last < a[j].Last
}


func main () {
    u1 := user {
        Last : "James",
        Age : 32,
        Sayings : []string{"B", "A"},
    }
    u2 := user {
        Last : "Moneypenny",
        Age : 27,
        Sayings : []string{"C", "B"},
    }
    u3 := user {
        Last : "Ian",
        Age : 34,
        Sayings : []string{"X", "Y"},
    }

    users := []user{u1, u2, u3}

    for _, v := range(users) {
            sort.Strings(v.Sayings)
    }

    sort.Sort(ByAge(users))
    fmt.Println(users)
    sort.Sort(ByLast(users))
    fmt.Println(users)

}
