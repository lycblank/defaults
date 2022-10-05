package json

import (
    libjson "encoding/json"
    "github.com/lycblank/defaults"
    "io"
)

type Decoder struct {
    *libjson.Decoder
}

// NewDecoder returns a new decoder that reads from r.
//
// The decoder introduces its own buffering and may
// read data from r beyond the JSON values requested.
func NewDecoder(r io.Reader) *Decoder {
    return &Decoder{
        Decoder: libjson.NewDecoder(r),
    }
}

// Decode reads the next JSON-encoded value from its
// input and stores it in the value pointed to by v.
//
// See the documentation for Unmarshal for details about
// the conversion of JSON into a Go value.
func (dec *Decoder) Decode(v interface{}) error {
    if err := dec.Decoder.Decode(v); err != nil {
        return err
    }
    return defaults.Apply(v)
}
