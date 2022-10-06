package main

import (
    "fmt"
    "github.com/gin-gonic/gin/binding"
    "github.com/lycblank/defaults/validator"
)

func main() {
    validator := validator.NewStructValidatorWithDefault(binding.Validator)
    // Access gin to support the default tag
    // binding.Validator = validator.NewStructValidatorWithDefault(binding.Validator)
    carWithDefault := &CarWithDefault{}
    if err := validator.ValidateStruct(carWithDefault); err != nil {
        panic(err)
    }
    fmt.Printf("car:%+v\n", carWithDefault) //Prints: car:&{Color:red Size:100 IsImport:true WheelRadius:1.01 SeatCount:4}

    carWithoutDefault := &CarWithoutDefault{}
    if err := validator.ValidateStruct(carWithoutDefault); err != nil {
        panic(err) // panic: Key: 'CarWithoutDefault.Color' Error:Field validation for 'Color' failed on the 'required' tag
    }
}

type CarWithDefault struct {
    Color       string  `default:"red" binding:"required"`
    Size        int     `default:"100" binding:"required"`
    IsImport    bool    `default:"true" binding:"required"`
    WheelRadius float64 `default:"1.01" binding:"required"`
    SeatCount   uint    `default:"4" binding:"required"`
}

type CarWithoutDefault struct {
    Color       string  `binding:"required"`
    Size        int     `binding:"required"`
    IsImport    bool    `binding:"required"`
    WheelRadius float64 `binding:"required"`
    SeatCount   uint    `binding:"required"`
}
