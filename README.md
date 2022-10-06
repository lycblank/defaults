defaults
============
Enabling struct with defaults values

Installation
-----
The recommended way to install defaults
```shell
go get -u -v github.com/lycblank/defaults
```

Examples
------
Set the default value of the structure field through the structure tag
```go
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
    fmt.Printf("car:%+v\n", car) //Prints: car:&{Color:red Size:100 IsImport:true WheelRadius:1.01 SeatCount:4}
}

type Car struct {
    Color       string  `default:"red"`
    Size        int     `default:"100"`
    IsImport    bool    `default:"true"`
    WheelRadius float64 `default:"1.01"`
    SeatCount uint `default:"4"`
}
```

Set the default value of the structure through json parsing
```go
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
```

Access gin to support the default tag
```go
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
```

Caveats
------
A value of zero will be overwritten as the default value if the default tag is set.

License
-------
MIT, see [LICENSE](https://github.com/lycblank/defaults/blob/master/LICENSE)
