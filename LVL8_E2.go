package main

import (
    "fmt"
    "encoding/json"
)

type user struct {
    First string
    Age int
}

func main () {
    u1 := user {
        First : "James",
        Age : 32,
    }
    u2 := user {
        First : "Moneypenny",
        Age : 27,
    }
    u3 := user {
        First : "Ian",
        Age : 34,
    }

    users := []user{u1, u2, u3}

    bs, err := json.Marshal(users)
    if err != nil {
        fmt.Println("err:", err)
    }
    fmt.Println(string(bs))

    usersUnmarshal := []user{u1, u2, u3}

    err = json.Unmarshal(bs, &usersUnmarshal)
    if err != nil {
        fmt.Println("err:", err)
    }
    fmt.Println(usersUnmarshal)
}
