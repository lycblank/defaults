package json

import (
    libjson "encoding/json"
    "io"
)

// NewEncoder returns a new encoder that writes to w.
func NewEncoder(w io.Writer) *libjson.Encoder {
    return libjson.NewEncoder(w)
}
