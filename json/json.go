package json

import (
    libjson "encoding/json"
    "github.com/lycblank/defaults"
)

func Unmarshal(data []byte, v interface{}) error {
    if err := libjson.Unmarshal(data, v); err != nil {
        return err
    }

    return defaults.Apply(v)
}

func Marshal(v interface{}) ([]byte, error) {
    return libjson.Marshal(v)
}

