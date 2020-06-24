package main

import (
    "fmt"
    "log"
    "errors"
)

type sqrtError struct {
    lat string
    long string
    err error
}

func (se sqrtError) Error() string {
    return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}


func main () {
    _, err := sqrt(-10.23)
    if err != nil {
        log.Println(err)
    }
}

func sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, sqrtError{
            lat : "3.108",
            long : "2.781",
            err : errors.New("sqrt of negative is undefined"),
        }
    }
    return 42, nil
}
