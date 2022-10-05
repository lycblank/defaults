package defaults

import (
    "fmt"
    "reflect"
    "strconv"
)

type field struct {
    v reflect.Value
}

// check if the field is empty
func (f *field) isEmpty() bool {
    return f.v.IsZero()
}

// set value to field
func (f *field) set(val string) error {
    if !f.v.CanSet() || !f.v.CanAddr() {
        return ValueCanNotSet
    }
    switch f.v.Kind() {
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        v, err := strconv.ParseInt(val, 10, 64)
        if err != nil {
            return fmt.Errorf("can not convert %s to int64. field:%s err:%w", val, f.v.Type().Name(), err)
        }
        f.v.SetInt(v)
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
        v, err := strconv.ParseUint(val, 10, 64)
        if err != nil {
            return fmt.Errorf("can not convert %s to uint64. field:%s err:%w", val, f.v.Type().Name(), err)
        }
        f.v.SetUint(v)
    case reflect.Float32, reflect.Float64:
        v, err := strconv.ParseFloat(val, 10)
        if err != nil {
            return fmt.Errorf("can not convert %s to float64. field:%s err:%w", val, f.v.Type().Name(), err)
        }
        f.v.SetFloat(v)
    case reflect.String:
        f.v.SetString(val)
    case reflect.Bool:
        v, err := strconv.ParseBool(val)
        if err != nil {
            return fmt.Errorf("can not convert %s to bool. field:%s err:%w", val, f.v.Type().Name(), err)
        }
        f.v.SetBool(v)
    default:
        return DataTypeNotSupport
    }
    return nil
}

// set default value to filed
func (f *field) apply(tag reflect.StructTag) error {
    if !f.isEmpty() {
        return nil
    }
    if !f.v.CanSet() {
        return FieldCanNotSet
    }

    defaultValue := tag.Get(defaultTag)
    return f.set(defaultValue)
}
