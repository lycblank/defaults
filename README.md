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
Caveats
------
A value of zero will be overwritten as the default value if the default tag is set.

License
-------
MIT, see [LICENSE](https://github.com/lycblank/defaults/blob/master/LICENSE)
