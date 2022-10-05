package defaults

import "reflect"

// Apply set default values for variables
func Apply(v interface{}) error {
    rv := reflect.ValueOf(v)
    if rv.Kind() != reflect.Ptr {
        return ApplyValueNotPtr
    }
    rv = rv.Elem()

    if rv.Kind() == reflect.Struct {
        return applyStruct(rv)
    }

    return DataTypeNotSupport
}

// set default values for struct
func applyStruct(v reflect.Value) error {
    field := &field{}

    numField := v.NumField()
    t := v.Type()
    for i := 0; i < numField; i++ {
        field.v = v.Field(i)
        if err := field.apply(t.Field(i).Tag); err != nil {
            return err
        }
    }
    return nil
}



