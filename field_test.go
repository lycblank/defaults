package defaults

import (
    "github.com/stretchr/testify/assert"
    "reflect"
    "strconv"
    "testing"
)

func TestIsEmpty(t *testing.T) {
    type testCase struct {
        v      interface{}
        msg    string
        expect interface{}
    }
    cases := []testCase{
        {v: "", msg: "string should be empty", expect: true},
        {v: int(0), msg: "int should be empty", expect: true},
        {v: float64(0), msg: "float64 should be empty", expect: true},
        {v: false, msg: "bool should be empty", expect: true},
        {v: uint(0), msg: "uint should be empty", expect: true},
        {v: "1", msg: "string should not be empty", expect: false},
        {v: int(1), msg: "int should not be empty", expect: false},
        {v: float64(1), msg: "float64 not should be empty", expect: false},
        {v: true, msg: "bool should not be empty", expect: false},
        {v: uint(1), msg: "uint should not be empty", expect: false},
    }
    for _, cas := range cases {
        f := &field{
            v: reflect.ValueOf(cas.v),
        }
        assert.Equal(t, cas.expect, f.isEmpty(), cas.msg)
    }
}

func TestSet(t *testing.T) {
    type testCase struct {
        v     int
        set    string
        msg    string
        exceptHasError bool
    }
    cases := []testCase{
        {v: int(0), set: "1", msg: "int should be 1", exceptHasError: false},
        {v: int(0), set: "xxxx", msg: "int should be 1", exceptHasError: true},
    }
    for _, cas := range cases {
        f := &field{
            v: reflect.ValueOf(&cas.v).Elem(),
        }

        assert.Equal(t, cas.exceptHasError, f.set(cas.set)!=nil, cas.msg)
        if !cas.exceptHasError {
            setV, _ := strconv.Atoi(cas.set)
            assert.Equal(t, cas.v, setV, cas.msg)
        }
    }
}
