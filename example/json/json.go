package main

import (
    "fmt"
    "github.com/lycblank/defaults/json"
)

func main() {
    raw := `{}`
    car := &Car{}
    if err := json.Unmarshal([]byte(raw), car); err != nil {
        panic(err)
    }
    fmt.Printf("car:%+v\n", car) //Prints: car:&{Color:red Size:100 IsImport:true WheelRadius:1.01 SeatCount:4}
}

type Car struct {
    Color       string  `default:"red"`
    Size        int     `default:"100"`
    IsImport    bool    `default:"true"`
    WheelRadius float64 `default:"1.01"`
    SeatCount   uint    `default:"4"`
}
