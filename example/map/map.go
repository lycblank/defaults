package main

import (
    "fmt"
    "github.com/lycblank/defaults"
)

func main() {
    car := &Car{Wheels: map[string]*CarWheel{
        "1": {},
        "2": {},
        "3": {},
        "4": {Radius: 5},
    }}
    if err := defaults.Apply(car); err != nil {
        panic(err)
    }
    for k, v := range car.Wheels {
        fmt.Printf("%s=%+v\n", k, v)
        // Prints:
        //1=&{Color:black Radius:1.02}
        //2=&{Color:black Radius:1.02}
        //3=&{Color:black Radius:1.02}
        //4=&{Color:black Radius:5}
    }
}

type Car struct {
    Color     string               `default:"red"`
    Size      int                  `default:"100"`
    IsImport  bool                 `default:"true"`
    SeatCount uint                 `default:"4"`
    Wheels    map[string]*CarWheel // note: The value of the map needs to be a pointer type
}

type CarWheel struct {
    Color  string  `default:"black"`
    Radius float64 `default:"1.02"`
}
