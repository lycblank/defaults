package main

import (
    "fmt"
    "github.com/lycblank/defaults"
)

func main() {
    car := &Car{Wheel: &CarWheel{}}
    if err := defaults.Apply(car); err != nil {
        panic(err)
    }
    fmt.Printf("car:%+v wheel：%+v\n", car, car.Wheel) //Prints: car:&{Color:red Size:100 IsImport:true SeatCount:4 Wheel:0xc00000c048} wheel：&{Color:black Radius:1.02}
}

type Car struct {
    Color       string  `default:"red"`
    Size        int     `default:"100"`
    IsImport    bool    `default:"true"`
    SeatCount   uint    `default:"4"`
    Wheel *CarWheel
}

type CarWheel struct {
    Color string `default:"black"`
    Radius float64 `default:"1.02"`
}