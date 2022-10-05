package defaults

import "errors"

var (
    defaultTag = "default"
)

var (
    FieldCanNotSet = errors.New("field can not set")
    ApplyValueNotPtr = errors.New("apply value not ptr")
    DataTypeNotSupport = errors.New("data type not support")
)
