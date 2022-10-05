package main

import (
    "fmt"
    "github.com/lycblank/defaults"
)

func main() {
    car := &Car{}
    if err := defaults.Apply(car); err != nil {
        panic(err)
    }
    fmt.Printf("car:%+v\n", car)
}

type Car struct {
    Color       string  `default:"red"`
    Size        int     `default:"100"`
    IsImport    bool    `default:"true"`
    WheelRadius float64 `default:"1.01"`
    SeatCount uint `default:"4"`
}
